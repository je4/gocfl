package extension

import (
	"bytes"
	"crypto/tls"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/andybalholm/brotli"
	"github.com/je4/gocfl/v2/pkg/ocfl"
	ironmaiden "github.com/je4/indexer/pkg/indexer"
	"io"
	"net/http"
	"net/url"
)

const IndexerName = "NNNN-indexer"
const IndexerDescription = "technical metadata for all files"

func GetIndexerParams() []*ocfl.ExtensionExternalParam {
	return []*ocfl.ExtensionExternalParam{
		{
			ExtensionName: IndexerName,
			Param:         "indexer-url",
			File:          "IndexerUrl",
			Description:   "url for indexer format recognition service",
		},
	}
}

type IndexerConfig struct {
	*ocfl.ExtensionConfig
	StorageType string
	StorageName string
}
type Indexer struct {
	*IndexerConfig
	fs         ocfl.OCFLFS
	indexerURL *url.URL
	sourceFS   ocfl.OCFLFSRead
	buffer     *bytes.Buffer
	writer     *brotli.Writer
}

func NewIndexerFS(fs ocfl.OCFLFSRead, urlString string, sourceFS ocfl.OCFLFSRead) (*Indexer, error) {
	fp, err := fs.Open("config.json")
	if err != nil {
		return nil, errors.Wrap(err, "cannot open config.json")
	}
	defer fp.Close()
	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read config.json")
	}

	var config = &IndexerConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal DirectCleanConfig '%s'", string(data))
	}
	return NewIndexer(config, urlString, sourceFS)
}
func NewIndexer(config *IndexerConfig, urlString string, sourceFS ocfl.OCFLFSRead) (*Indexer, error) {
	var err error
	sl := &Indexer{
		IndexerConfig: config,
		sourceFS:      sourceFS,
		buffer:        new(bytes.Buffer),
	}
	sl.writer = brotli.NewWriter(sl.buffer)
	if sl.indexerURL, err = url.Parse(urlString); err != nil {
		return nil, errors.Wrapf(err, "cannot parse url '%s'", urlString)
	}
	if config.ExtensionName != sl.GetName() {
		return nil, errors.New(fmt.Sprintf("invalid extension name'%s'for extension %s", config.ExtensionName, sl.GetName()))
	}
	return sl, nil
}

func (sl *Indexer) GetConfigString() string {
	str, _ := json.MarshalIndent(sl.IndexerConfig, "", "  ")
	return string(str)
}

func (sl *Indexer) IsRegistered() bool { return false }

func (sl *Indexer) GetName() string { return IndexerName }

func (sl *Indexer) SetFS(fs ocfl.OCFLFS) { sl.fs = fs }

func (sl *Indexer) SetParams(params map[string]string) error {
	var err error
	name := fmt.Sprintf("ext-%s-%s", IndexerName, "indexer-url")
	urlString, _ := params[name]
	if urlString == "" {
		if sl.indexerURL != nil {
			result, code, err := sl.post("{}")
			if err != nil {
				return errors.Wrapf(err, "cannot post to '%s'", urlString)
			}
			if code != http.StatusBadRequest {
				return errors.Errorf("cannot post to '%s' - %v:'%s'", urlString, code, result)
			}
			_ = result
			return nil
		}
		return errors.Errorf("url '%s' not set", name)
	}
	if sl.indexerURL, err = url.Parse(urlString); err != nil {
		return errors.Wrapf(err, "cannot parse '%s' '%s'", name, urlString)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	result, code, err := sl.post("")
	if err != nil {
		return errors.Wrapf(err, "cannot post to '%s'", urlString)
	}
	if code != http.StatusBadRequest {
		return errors.Errorf("cannot post to '%s' - %v:'%s'", urlString, code, result)
	}
	_ = result

	return nil
}

func (sl *Indexer) post(data any) ([]byte, int, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot marshal %v", data)
	}
	resp, err := http.Post(sl.indexerURL.String(), "test/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot post %v to %s", data, sl.indexerURL)
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot read result of post %v to %s", data, sl.indexerURL)
	}
	return result, resp.StatusCode, nil
}

func (sl *Indexer) WriteConfig() error {
	if sl.fs == nil {
		return errors.New("no filesystem set")
	}
	configWriter, err := sl.fs.Create("config.json")
	if err != nil {
		return errors.Wrap(err, "cannot open config.json")
	}
	defer configWriter.Close()
	jenc := json.NewEncoder(configWriter)
	jenc.SetIndent("", "   ")
	if err := jenc.Encode(sl.ExtensionConfig); err != nil {
		return errors.Wrapf(err, "cannot encode config to file")
	}
	return nil
}

func (sl *Indexer) AddFileBefore(object ocfl.Object, source, dest string) error {
	return nil
}
func (sl *Indexer) UpdateFileBefore(object ocfl.Object, source, dest string) error {
	return nil
}
func (sl *Indexer) DeleteFileBefore(object ocfl.Object, dest string) error {
	// nothing to do
	return nil
}
func (sl *Indexer) AddFileAfter(object ocfl.Object, source, internalPath, digest string) error {
	filePath := fmt.Sprintf("%s/%s", sl.sourceFS.String(), source)
	param := ironmaiden.ActionParam{
		Url:        filePath,
		Actions:    []string{"siegfried", "ffprobe", "identify"},
		HeaderSize: 0,
		Checksums:  map[string]string{},
	}
	result, code, err := sl.post(param)
	if err != nil {
		return errors.Wrapf(err, "indexer error for '%s'", filePath)
	}
	if code >= 300 {
		return errors.Errorf("indexer error for '%s': %s", filePath, result)
	}
	var meta = map[string]any{}
	if err := json.Unmarshal(result, &meta); err != nil {
		return errors.Errorf("cannot unmarshal indexer result `%s`", string(result))
	}
	var indexerline = struct {
		Digest   string
		Metadata map[string]any
	}{
		Digest:   digest,
		Metadata: meta,
	}
	data, err := json.Marshal(indexerline)
	if err != nil {
		return errors.Errorf("cannot marshal result %v", indexerline)
	}
	if _, err := sl.writer.Write(data); err != nil {
		return errors.Errorf("cannot brotli %s", string(data))
	}
	if _, err := sl.writer.Write([]byte("\n")); err != nil {
		return errors.Errorf("cannot brotli %s", string(data))
	}
	return nil
}
func (sl *Indexer) UpdateFileAfter(object ocfl.Object, source, dest string) error {
	return nil
}
func (sl *Indexer) DeleteFileAfter(object ocfl.Object, dest string) error {
	// nothing to do
	return nil
}

func (sl *Indexer) UpdateObjectBefore(object ocfl.Object) error {
	return nil
}

func (sl *Indexer) UpdateObjectAfter(object ocfl.Object) error {
	if err := sl.writer.Flush(); err != nil {
		return errors.Wrap(err, "cannot flush brotli writer")
	}
	if err := sl.writer.Close(); err != nil {
		return errors.Wrap(err, "cannot close brotli writer")
	}
	reader := brotli.NewReader(sl.buffer)
	switch sl.StorageType {
	case "area":
	case "path":
		//sl.buffer.Reset()
		targetname := fmt.Sprintf("%s/indexer_%s.jsonl", sl.IndexerConfig.StorageName, object.GetInventory().GetHead())
		if err := object.AddReader(io.NopCloser(reader), targetname, ""); err != nil {
			return errors.Wrapf(err, "cannot write '%s'", targetname)
		}
	case "extension":
	default:
		return errors.Errorf("unsupported storage type '%s'", sl.StorageType)
	}
	return nil
}

var (
	_ ocfl.Extension              = &Indexer{}
	_ ocfl.ExtensionContentChange = &Indexer{}
	_ ocfl.ExtensionObjectChange  = &Indexer{}
)
