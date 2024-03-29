package cmd

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/je4/filesystem/v2/pkg/writefs"
	"github.com/je4/gocfl/v2/pkg/ocfl"
	lm "github.com/je4/utils/v2/pkg/logger"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var extractMetaCmd = &cobra.Command{
	Use:     "extractmeta [path to ocfl structure]",
	Aliases: []string{},
	Short:   "extract metadata from ocfl structure",
	//Long:    "an utterly useless command for testing",
	Example: "gocfl extractmeta ./archive.zip --output-json ./archive_meta.json",
	Args:    cobra.ExactArgs(1),
	Run:     doExtractMeta,
}

func initExtractMeta() {
	extractMetaCmd.Flags().StringP("object-path", "p", "", "object path to extract")
	extractMetaCmd.Flags().StringP("object-id", "i", "", "object id to extract")
	extractMetaCmd.Flags().String("version", "latest", "version to extract")
	extractMetaCmd.Flags().String("format", "json", "output format (json)")
	extractMetaCmd.Flags().String("output", "", "output file (default stdout)")
}

func doExtractMetaConf(cmd *cobra.Command) {
	if str := getFlagString(cmd, "object-path"); str != "" {
		conf.ExtractMeta.ObjectPath = str
	}
	if str := getFlagString(cmd, "object-id"); str != "" {
		conf.ExtractMeta.ObjectID = str
	}
	if str := getFlagString(cmd, "version"); str != "" {
		conf.ExtractMeta.Version = str
	}
	if conf.ExtractMeta.Version == "" {
		conf.ExtractMeta.Version = "latest"
	}
	if str := getFlagString(cmd, "format"); str != "" {
		conf.ExtractMeta.Format = str
	}
	if str := getFlagString(cmd, "output"); str != "" {
		conf.ExtractMeta.Output = str
	}
}

func doExtractMeta(cmd *cobra.Command, args []string) {
	ocflPath, err := ocfl.Fullpath(args[0])
	if err != nil {
		cobra.CheckErr(err)
		return
	}

	daLogger, lf := lm.CreateLogger("ocfl", persistentFlagLogfile, nil, conf.LogLevel, conf.LogFormat)
	defer lf.Close()
	t := startTimer()
	defer func() { daLogger.Infof("Duration: %s", t.String()) }()

	doExtractMetaConf(cmd)

	oPath := conf.ExtractMeta.ObjectPath
	oID := conf.ExtractMeta.ObjectID
	if oPath != "" && oID != "" {
		cmd.Help()
		cobra.CheckErr(errors.New("do not use object-path AND object-id at the same time"))
		return
	}
	format := strings.ToLower(conf.ExtractMeta.Format)
	if format != "json" {
		cmd.Help()
		cobra.CheckErr(errors.Errorf("invalid format '%s' for flag 'format' or 'Format' config file entry", format))
		return
	}
	output := conf.ExtractMeta.Output

	daLogger.Infof("extracting metadata from '%s'", ocflPath)

	fsFactory, err := initializeFSFactory(nil, nil, nil, true, true, daLogger)
	if err != nil {
		daLogger.Errorf("cannot create filesystem factory: %v", err)
		daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		return
	}

	ocflFS, err := fsFactory.Get(ocflPath)
	if err != nil {
		daLogger.Errorf("cannot get filesystem for '%s': %v", ocflPath, err)
		daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		return
	}
	defer func() {
		if err := writefs.Close(ocflFS); err != nil {
			daLogger.Errorf("cannot close filesystem: %v", err)
			daLogger.Errorf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		}
	}()

	extensionParams := GetExtensionParamValues(cmd, conf)
	extensionFactory, err := InitExtensionFactory(extensionParams, "", false, nil, nil, nil, nil, daLogger)
	if err != nil {
		daLogger.Errorf("cannot initialize extension factory: %v", err)
		daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		return
	}

	ctx := ocfl.NewContextValidation(context.TODO())
	storageRoot, err := ocfl.LoadStorageRoot(ctx, ocflFS, extensionFactory, daLogger)
	if err != nil {
		daLogger.Errorf("cannot open storage root: %v", err)
		daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		return
	}

	metadata, err := storageRoot.ExtractMeta(oPath, oID)
	if err != nil {
		fmt.Printf("cannot extract metadata from storage root: %v\n", err)
		daLogger.Errorf("cannot extract metadata from storage root: %v\n", err)
		daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		return
	}

	jsonBytes, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		fmt.Printf("cannot marshal metadata")
		daLogger.Errorf("cannot marshal metadata: %v\n", err)
		daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
		return
	}
	if output != "" {
		if err := os.WriteFile(output, jsonBytes, 0644); err != nil {
			fmt.Printf("cannot write json to file")
			daLogger.Errorf("cannot write json to file '%s': %v\n", output, err)
			daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
			return
		}
	} else {
		if _, err := os.Stdout.Write(jsonBytes); err != nil {
			fmt.Printf("cannot write json to file")
			daLogger.Errorf("cannot write json to file standard output: %v\n", err)
			daLogger.Debugf("%v%+v", err, ocfl.GetErrorStacktrace(err))
			return
		}
		fmt.Print("\n")
	}
	fmt.Printf("metadata extraction done without errors\n")
	showStatus(ctx)
}
