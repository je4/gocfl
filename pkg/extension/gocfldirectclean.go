package extension

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"go.ub.unibas.ch/gocfl/v2/pkg/ocfl"
	"io"
	"regexp"
	"strings"
)

const DirectCleanName = "NNNN-direct-clean-path-layout"

var directCleanRuleAll = regexp.MustCompile("[\u0000-\u001f\u007f\u0020\u0085\u00a0\u1680\u2000-\u200f\u2028\u2029\u202f\u205f\u3000\n\t*?:\\[\\]\"<>|(){}&'!\\;#@]")
var directCleanRuleWhitespace = regexp.MustCompile("[\u0009\u000a-\u000d\u0020\u0085\u00a0\u1680\u2000-\u200f\u2028\u2029\u202f\u205f\u3000]")
var directCleanRule_1_5 = regexp.MustCompile("[\u0000-\u001F\u007F\n\r\t*?:\\[\\]\"<>|(){}&'!\\;#@]")
var directCleanRule_2_4_6 = regexp.MustCompile("^[\\-~\u0009\u000a-\u000d\u0020\u0085\u00a0\u1680\u2000-\u200f\u2028\u2029\u202f\u205f\u3000]*(.*?)[\u0009\u000a-\u000d\u0020\u0085\u00a0\u1680\u2000-\u20a0\u2028\u2029\u202f\u205f\u3000]*$")
var directCleanRulePeriods = regexp.MustCompile("^\\.+$")

var directCleanErrFilenameTooLong = errors.New("filename too long")
var directCleanErrPathnameTooLong = errors.New("pathname too long")

type DirectClean struct {
	*DirectCleanConfig
}

type DirectCleanConfig struct {
	*ocfl.ExtensionConfig
	MaxPathnameLen              int    `json:"maxPathnameLen,omitempty"`
	MaxFilenameLen              int    `json:"maxFilenameLen,omitempty"`
	ReplacementString           string `json:"replacementString,omitempty"`
	WhitespaceReplacementString string `json:"whitespaceReplacementString,omitempty"`
	UTFEncode                   bool   `json:"utfEncode,omitempty"`
}

func NewDirectCleanFS(fs ocfl.OCFLFS) (ocfl.Extension, error) {
	fp, err := fs.Open("config.json")
	if err != nil {
		return nil, errors.Wrap(err, "cannot open config.json")
	}
	defer fp.Close()
	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read config.json")
	}
	var config = &DirectCleanConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal DirectCleanConfig '%s'", string(data))
	}
	return NewDirectClean(config)
}

func NewDirectClean(config *DirectCleanConfig) (ocfl.Extension, error) {
	if config.MaxPathnameLen == 0 {
		config.MaxPathnameLen = 32000
	}
	if config.MaxFilenameLen == 0 {
		config.MaxFilenameLen = 127
	}
	sl := &DirectClean{DirectCleanConfig: config}
	if config.ExtensionName != sl.GetName() {
		return nil, errors.Errorf("invalid extension name %s for extension %s", config.ExtensionName, sl.GetName())
	}
	return sl, nil
}

func encodeUTFCode(s string) string {
	return "=u" + strings.Trim(fmt.Sprintf("%U", []rune(s)), "U+[]")
}

// interface Extension
func (*DirectClean) IsObjectExtension() bool      { return true }
func (*DirectClean) IsStoragerootExtension() bool { return true }

func (sl *DirectClean) GetName() string { return DirectCleanName }

// interface Extension
func (sl *DirectClean) WriteConfig(configWriter io.Writer) error {
	jenc := json.NewEncoder(configWriter)
	jenc.SetIndent("", "   ")
	if err := jenc.Encode(sl.DirectCleanConfig); err != nil {
		return errors.Wrapf(err, "cannot encode config to file")
	}
	return nil
}

// interface
func (sl *DirectClean) BuildStorageRootPath(storageRoot ocfl.StorageRoot, id string) (string, error) {
	return sl.build(id)
}

func (sl *DirectClean) BuildObjectContentPath(storageRoot ocfl.StorageRoot, id string) (string, error) {
	return sl.build(id)
}

func (sl *DirectClean) build(fname string) (string, error) {

	fname = strings.ToValidUTF8(fname, sl.ReplacementString)

	names := strings.Split(fname, "/")
	result := []string{}

	for _, n := range names {
		if len(n) == 0 {
			continue
		}
		if sl.UTFEncode {
			n = directCleanRuleAll.ReplaceAllStringFunc(n, encodeUTFCode)
			if n[0] == '~' || directCleanRulePeriods.MatchString(n) {
				n = encodeUTFCode(string(n[0])) + n[1:]
			}
		} else {
			n = directCleanRuleWhitespace.ReplaceAllString(n, sl.WhitespaceReplacementString)
			n = directCleanRule_1_5.ReplaceAllString(n, sl.ReplacementString)
			n = directCleanRule_2_4_6.ReplaceAllString(n, "$1")
			if directCleanRulePeriods.MatchString(n) {
				n = sl.ReplacementString + n[1:]
			}
		}

		lenN := len(n)
		if lenN > sl.MaxFilenameLen {
			return "", errors.Wrapf(directCleanErrFilenameTooLong, "filename: %s", n)
		}

		if lenN > 0 {
			result = append(result, n)
		}
	}

	fname = strings.Join(result, "/")

	if len(fname) > sl.MaxPathnameLen {
		return "", errors.Wrapf(directCleanErrPathnameTooLong, "pathname: %s", fname)
	}

	return fname, nil
}