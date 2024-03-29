package ocfl

import (
	"emperror.dev/errors"
	"encoding/json"
	"golang.org/x/exp/slices"
	"time"
)

type InventorySpec string

const (
	InventorySpec1_0 InventorySpec = "https://ocfl.io/1.0/spec/#inventory"
	InventorySpec1_1 InventorySpec = "https://ocfl.io/1.1/spec/#inventory"
)

// return true if Specification s1 < s2
func SpecIsLessOrEqual(s1, s2 InventorySpec) bool {
	//return s1 == InventorySpec1_0 && s2 == InventorySpec1_1
	return s1 <= s2
}

type OCFLState struct {
	State map[string][]string
	err   error
}

func (s *OCFLState) UnmarshalJSON(data []byte) error {
	s.State = map[string][]string{}
	if err := json.Unmarshal(data, &s.State); err != nil {
		s.err = errors.Wrapf(err, "cannot unmarshal state %s", string(data))
		return nil
	}
	return nil
}

func (s *OCFLState) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.State)
}

type OCFLString struct {
	string
	err error
}

func NewOCFLString(str string) *OCFLString {
	return &OCFLString{
		string: str,
		err:    nil,
	}
}

func (s *OCFLString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		s.err = errors.Wrapf(err, "cannot unmarshal string '%s'", string(data))
		return nil
	}
	s.string = str
	return nil
}

func (s *OCFLString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.string)
}

func (s *OCFLString) String() string {
	return s.string
}

type OCFLTime struct {
	time.Time
	err error
}

func (t *OCFLTime) MarshalJSON() ([]byte, error) {
	tstr := t.Format(time.RFC3339)
	return json.Marshal(tstr)
}
func (t *OCFLTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		t.err = errors.Wrapf(err, "cannot unmarshal time '%s'", string(data))
		return nil
	}
	tt, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.err = errors.Wrapf(err, "cannot parse time %s", string(data))
		return nil
	}
	t.Time = tt

	return nil
}

type User struct {
	Address *OCFLString `json:"address,omitempty"`
	Name    *OCFLString `json:"name"`
}
type OCFLUser struct {
	User
	err error
}

func NewOCFLUser(name, address string) *OCFLUser {
	user := &OCFLUser{
		User: User{
			Address: NewOCFLString(address),
			Name:    NewOCFLString(name),
		},
		err: nil,
	}
	return user
}

func (u *OCFLUser) UnmarshalJSON(data []byte) error {
	tu := &User{}
	if err := json.Unmarshal(data, tu); err != nil {
		u.err = errors.Wrapf(err, "cannot unmarshal user '%s'", string(data))
		return nil
	}
	u.User.Address = tu.Address
	u.User.Name = tu.Name

	return nil
}

func (u *OCFLUser) Finalize() {
	if u.Name == nil {
		u.Name = NewOCFLString("")
	}
	if u.Address == nil {
		u.Address = NewOCFLString("")
	}
}

type Version struct {
	Created *OCFLTime   `json:"created"`
	Message *OCFLString `json:"message"`
	State   *OCFLState  `json:"state"`
	User    *OCFLUser   `json:"user"`
}

func (v *Version) EqualMeta(v2 *Version) bool {
	if v2 == nil {
		return false
	}
	if v.Created.Time.String() != v2.Created.Time.String() ||
		v.Message.string != v2.Message.string ||
		v.User.Name.string != v2.User.Name.string ||
		v.User.Address.string != v2.User.Address.string {
		return false
	}
	return true
}
func (v *Version) EqualState(v2 *Version) bool {
	if v2 == nil {
		return false
	}
	/*
		if v.Created.Time.String() != v2.Created.Time.String() ||
			v.Message.string != v2.Message.string ||
			v.User.GetName.string != v2.User.GetName.string ||
			v.User.Address.string != v2.User.Address.string {
			return false
		}
	*/
	if len(v.State.State) != len(v2.State.State) {
		return false
	}
	files := []string{}
	for _, vals := range v.State.State {
		files = append(files, vals...)
	}
	slices.Sort(files)
	files = slices.Compact(files)
	files2 := []string{}
	for _, vals := range v2.State.State {
		files2 = append(files2, vals...)
	}
	slices.Sort(files2)
	files2 = slices.Compact(files2)
	if len(files) != len(files2) {
		return false
	}
	if !sliceContains(files, files2) {
		return false
	}
	return true
}

type OCFLManifest struct {
	Manifest map[string][]string
	err      error
}

func (m *OCFLManifest) UnmarshalJSON(data []byte) error {
	m.Manifest = map[string][]string{}
	if err := json.Unmarshal(data, &m.Manifest); err != nil {
		m.err = errors.Wrapf(err, "cannot unmarshal versions '%s'", string(data))
		return nil
	}

	return nil
}

func (m *OCFLManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Manifest)
}

type OCFLVersions struct {
	Versions map[string]*Version
	err      error
}

func (v *OCFLVersions) GetVersion(version string) (*Version, error) {
	ver, ok := v.Versions[version]
	if !ok {
		return nil, errors.Errorf("invalid version '%s'", version)
	}
	return ver, nil
}

func (v *OCFLVersions) UnmarshalJSON(data []byte) error {
	v.Versions = map[string]*Version{}
	if err := json.Unmarshal(data, &v.Versions); err != nil {
		v.err = errors.Wrapf(err, "cannot unmarshal versions '%s'", string(data))
		return nil
	}

	return nil
}

func (v *OCFLVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Versions)
}
