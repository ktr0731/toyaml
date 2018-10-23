package toyaml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ToYAML receives TOML encoded bytes `t` and
// converts `t` to YAML encoded bytes.
// `t` must be []byte, string or io.Reader.
// If some fields have "toml" tag, its "json" tag's value must be same value "toml" one.
func ToYAML(t interface{}) ([]byte, error) {
	var m map[string]interface{}
	var err error
	switch t := t.(type) {
	case []byte:
		err = toml.Unmarshal(t, &m)
	case string:
		_, err = toml.Decode(t, &m)
	case io.Reader:
		_, err = toml.DecodeReader(t, &m)
	default:
		return nil, errors.Errorf("unsupported type: %T", t)
	}
	if err != nil {
		return nil, err
	}

	b, err := yaml.Marshal(m)
	if err != nil {
		return nil, err
	}

	return b, nil
}
