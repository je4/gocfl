package extension

import (
	"bufio"
	"bytes"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/andybalholm/brotli"
	"github.com/je4/gocfl/v2/pkg/migration"
	"github.com/je4/gocfl/v2/pkg/ocfl"
	"github.com/je4/indexer/v2/pkg/indexer"
	"golang.org/x/exp/slices"
	"io"
	"regexp"
)

const MigrationName = "NNNN-migration"
const MigrationDescription = "preservation management - file migration"

type MigrationConfig struct {
	*ocfl.ExtensionConfig
	StorageType string
	StorageName string
	Compress    string
}

type MigrationTarget struct {
	Name            string
	Strategy        string        // add or replace
	FilenameRegexp  regexp.Regexp // target expression
	FilenameReplace string        // replacement string
	Command         string        // command to execute (stdin --> stdout)
}

type migrationResult struct {
	Source string `json:"source,omitempty"`
	Error  string `json:"error,omitempty"`
	ID     string `json:"id"`
}

type migrationLine struct {
	Path      string           `json:"path"`
	Migration *migrationResult `json:"migration"`
}

// map pronom to migration
type MigrationMap map[string]*MigrationTarget

// map checksum to migration
type MigrationFiles map[string]*MigrationTarget

type Migration struct {
	*MigrationConfig
	fs        ocfl.OCFLFSRead
	lastHead  string
	migration *migration.Migration
	//buffer         *bytes.Buffer
	buffer         map[string]*bytes.Buffer
	writer         *brotli.Writer
	migrationFiles map[string]*migration.Function
	migratedFiles  map[string]map[string]string
	sourceFS       ocfl.OCFLFSRead
	currentHead    string
}

func NewMigrationFS(fs ocfl.OCFLFSRead, migration *migration.Migration) (*Migration, error) {
	fp, err := fs.Open("config.json")
	if err != nil {
		return nil, errors.Wrap(err, "cannot open config.json")
	}
	defer fp.Close()
	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read config.json")
	}

	var config = &MigrationConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal DirectCleanConfig '%s'", string(data))
	}
	ext, err := NewMigration(config, migration)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create new indexer")
	}
	return ext, nil
}
func NewMigration(config *MigrationConfig, mig *migration.Migration) (*Migration, error) {
	sl := &Migration{
		MigrationConfig: config,
		migration:       mig,
		buffer:          map[string]*bytes.Buffer{},
		migrationFiles:  map[string]*migration.Function{},
		migratedFiles:   map[string]map[string]string{},
	}
	//	sl.writer = brotli.NewWriter(sl.buffer)
	if config.ExtensionName != sl.GetName() {
		return nil, errors.New(fmt.Sprintf("invalid extension name'%s'for extension %s", config.ExtensionName, sl.GetName()))
	}
	if mig != nil {
		sl.sourceFS = mig.SourceFS
	}
	return sl, nil
}

func (mi *Migration) GetConfigString() string {
	str, _ := json.MarshalIndent(mi.MigrationConfig, "", "  ")
	return string(str)
}

func (mi *Migration) IsRegistered() bool { return false }

func (mi *Migration) GetName() string { return MigrationName }

func (mi *Migration) SetFS(fs ocfl.OCFLFSRead) { mi.fs = fs }

func (mi *Migration) SetParams(params map[string]string) error {
	return nil
}

func (mi *Migration) WriteConfig() error {
	if mi.fs == nil {
		return errors.New("no filesystem set")
	}
	fsRW, ok := mi.fs.(ocfl.OCFLFS)
	if !ok {
		return errors.Errorf("filesystem is read only - '%s'", mi.fs.String())
	}

	configWriter, err := fsRW.Create("config.json")
	if err != nil {
		return errors.Wrap(err, "cannot open config.json")
	}
	defer configWriter.Close()
	jenc := json.NewEncoder(configWriter)
	jenc.SetIndent("", "   ")
	if err := jenc.Encode(mi.MigrationConfig); err != nil {
		return errors.Wrapf(err, "cannot encode config to file")
	}
	return nil
}

func (mi *Migration) UpdateObjectBefore(object ocfl.Object) error {
	return nil
}

func (mi *Migration) alreadyMigrated(cs string) bool {
	for _, mf := range mi.migratedFiles {
		if _, ok := mf[cs]; ok {
			return true
		}
	}
	return false
}

func (mi *Migration) UpdateObjectAfter(object ocfl.Object) error {
	// first get the metadata from the object
	meta, err := object.GetMetadata()
	if err != nil {
		return errors.Wrapf(err, "cannot get metadata from object %s", object.GetID())
	}
	for cs, m := range meta.Files {
		indexerMetaAny, ok := m.Extension[IndexerName]
		if !ok {
			continue
		}
		indexerMeta, ok := indexerMetaAny.(*indexer.ResultV2)
		if !ok {
			continue
		}
		migration, err := mi.migration.GetFunctionByPronom(indexerMeta.Pronom)
		if err != nil {
			continue
			//return errors.Wrapf(err, "cannot get migration function for pronom %s", pronom)
		}
		if mi.alreadyMigrated(cs) {
			continue
		}
		if migrationMetaAny, ok := m.Extension[MigrationName]; ok {
			if migrationMetaMap, ok := migrationMetaAny.(map[string]any); ok {
				if _, ok := migrationMetaMap[cs]; ok {
					continue
				}
			}
		}

		mi.migrationFiles[cs] = migration
	}
	inventory := object.GetInventory()
	if inventory == nil {
		return errors.Errorf("inventory is nil")
	}
	mi.lastHead = inventory.GetHead()
	return nil
}

func (mi *Migration) NeedNewVersion(object ocfl.Object) (bool, error) {
	return len(mi.migrationFiles) > 0, nil
}

func (mi *Migration) DoNewVersion(object ocfl.Object) error {
	inventory := object.GetInventory()
	head := inventory.GetHead()
	extensionManager := object.GetExtensionManager()
	if extensionManager == nil {
		return errors.Errorf("extension manager is nil")
	}
	mi.buffer[head] = &bytes.Buffer{}
	mi.writer = brotli.NewWriter(mi.buffer[head])
	//files := inventory.GetFiles()

	versions := inventory.GetVersionStrings()
	if len(versions) < 2 {
		return errors.Errorf("cannot migrate files in object '%s' - no previous version", object.GetID())
	}
	manifest := inventory.GetManifest()
	if _, ok := mi.migratedFiles[head]; !ok {
		mi.migratedFiles[head] = map[string]string{}
	}
	for cs, mig := range mi.migrationFiles {
		// todo: do it more efficient
		var found = false
		for _, mf := range mi.migratedFiles {
			if _, ok := mf[cs]; ok {
				found = true
			}
		}
		if found {
			continue
		}

		var targetNames = []string{}
		manifestFiles, ok := manifest[cs]
		if !ok {
			return errors.Errorf("cannot find file with checksum '%s' in object '%s'", cs, object.GetID())
		}
		/*
			for _, f := range manifestFiles {
				targetNames = append(targetNames, mig.GetDestinationName(f))
			}
		*/
		// get the files from last version
		stateFiles, err := inventory.GetStateFiles(versions[len(versions)-2], cs)
		if err != nil {
			return errors.Wrapf(err, "cannot get state files for checksum '%s' in object '%s'", cs, object.GetID())
		}
		for _, sf := range stateFiles {
			targetNames = append(targetNames, mig.GetDestinationName(sf))
		}

		var file io.ReadCloser
		fs := object.GetFS()
		if fs != nil {
			file, err = fs.Open(manifestFiles[0])
			if err != nil {
				file = nil
			}
		}
		if file == nil {
			if mi.sourceFS != nil {
				stateFiles, err := inventory.GetStateFiles("", cs)
				if err != nil {
					return errors.Wrapf(err, "cannot get state files for checksum '%s' in object '%s'", cs, object.GetID())
				}
				if len(stateFiles) == 0 {
					return errors.Errorf("zero state file for checksum '%s' in object '%s'", cs, object.GetID())
				}
				external, err := object.GetExtensionManager().BuildObjectExternalPath(object, stateFiles[len(stateFiles)-1])
				if err != nil {
					return errors.Wrapf(err, "cannot build external path for file '%s' in object '%s'", stateFiles[len(stateFiles)-1], object.GetID())
				}
				file, err = mi.sourceFS.Open(external)
				if err != nil {
					return errors.Wrapf(err, "cannot open file '%s/%s' in source filesystem", mi.sourceFS.String(), targetNames[len(targetNames)-1])
				}
			}
		}
		var ml *migrationLine
		path, err := extensionManager.BuildObjectInternalPath(object, targetNames[0], "content")
		if err != nil {
			return errors.Wrapf(err, "cannot build state path for file '%s' in object '%s'", targetNames[0], object.GetID())
		}
		path = inventory.BuildManifestName(path)
		if err := migration.DoMigrate(object, mig, targetNames, file); err != nil {
			ml = &migrationLine{
				Path: path,
				Migration: &migrationResult{
					Source: manifestFiles[0],
					Error:  err.Error(),
					ID:     mig.GetID(),
				},
			}
			//			return err
		} else {
			ml = &migrationLine{
				Path: path,
				Migration: &migrationResult{
					Source: manifestFiles[0],
					ID:     mig.GetID(),
				},
			}
			switch mig.Strategy {
			case migration.StrategyReplace:
				for _, n := range stateFiles {
					if slices.Contains(targetNames, n) {
						continue
					}
					if err := object.DeleteFile(n, cs); err != nil {
						return errors.Wrapf(err, "cannot delete file '%s' in object '%s'", n, object.GetID())
					}
				}
			}
		}
		data, err := json.Marshal(ml)
		if err != nil {
			return errors.Wrapf(err, "cannot marshal migration line for file '%s' in object '%s'", targetNames[0], object.GetID())
		}
		if _, err := mi.writer.Write(append(data, []byte("\n")...)); err != nil {
			return errors.Wrapf(err, "cannot write migration line for file '%s' in object '%s'", targetNames[0], object.GetID())
		}
	}
	if err := mi.writer.Flush(); err != nil {
		return errors.Wrapf(err, "cannot flush migration line writer for object '%s'", object.GetID())
	}
	if err := mi.writer.Close(); err != nil {
		return errors.Wrapf(err, "cannot close migration line writer for object '%s'", object.GetID())
	}
	buffer, ok := mi.buffer[head]
	if !ok {
		return nil
	}
	if err := ocfl.WriteJsonL(
		object,
		"migration",
		buffer.Bytes(),
		mi.MigrationConfig.Compress,
		mi.StorageType,
		mi.StorageName,
		mi.fs,
	); err != nil {
		return errors.Wrap(err, "cannot write jsonl")
	}
	return nil
}

func (mi *Migration) GetMetadata(object ocfl.Object) (map[string]any, error) {
	var err error
	var result = map[string]any{}

	inventory := object.GetInventory()
	manifest := inventory.GetManifest()
	path2digest := map[string]string{}
	for checksum, names := range manifest {
		for _, name := range names {
			path2digest[name] = checksum
		}
	}
	for v, _ := range inventory.GetVersions() {
		var data []byte
		if buf, ok := mi.buffer[v]; ok && buf.Len() > 0 {
			//		if v == inventory.GetHead() && sl.buffer.Len() > 0 {
			// need a new reader on the buffer
			reader := brotli.NewReader(bytes.NewBuffer(buf.Bytes()))
			data, err = io.ReadAll(reader)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot read buffer for '%s' '%s'", object.GetID(), v)
			}
		} else {
			data, err = ocfl.ReadJsonL(object, "migration", v, mi.MigrationConfig.Compress, mi.StorageType, mi.StorageName, mi.fs)
			if err != nil {
				continue
				// return nil, errors.Wrapf(err, "cannot read jsonl for '%s' version '%s'", object.GetID(), v)
			}
		}

		reader := bytes.NewReader(data)
		r := bufio.NewScanner(reader)
		r.Buffer(make([]byte, 128*1024), 16*1024*1024)
		r.Split(bufio.ScanLines)
		for r.Scan() {
			line := r.Text()
			var meta = migrationLine{}
			if err := json.Unmarshal([]byte(line), &meta); err != nil {
				return nil, errors.Wrapf(err, "cannot unmarshal line from for '%s' %s - [%s]", object.GetID(), v, line)
			}
			var digest string
			for cs, names := range manifest {
				for _, name := range names {
					if name == meta.Migration.Source {
						digest = cs
						break
					}
				}
			}
			if digest == "" {
				return nil, errors.Errorf("cannot find checksum for file '%s' in object '%s'", meta.Migration.Source, object.GetID())
			}
			cs, ok := path2digest[meta.Path]
			if !ok && meta.Migration.Error != "" {
				cs, ok = path2digest[meta.Migration.Source]
			}
			if !ok {
				return nil, errors.Errorf("cannot find checksum for file '%s' in object '%s'", meta.Path, object.GetID())
			}
			meta.Migration.Source = digest
			result[cs] = meta.Migration
		}
		if err := r.Err(); err != nil {
			return nil, errors.Wrapf(err, "cannot scan lines for '%s' %s", object.GetID(), v)
		}
	}
	return result, nil
}

var (
	_ ocfl.Extension             = &Migration{}
	_ ocfl.ExtensionObjectChange = &Migration{}
	_ ocfl.ExtensionMetadata     = &Migration{}
	_ ocfl.ExtensionNewVersion   = &Migration{}
)
