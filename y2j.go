package y2j

import (
	json "github.com/rwxrob/json"
	"gopkg.in/yaml.v3"
)

// Convert converts YAML map data into JSON data. Only maps will work.
func Convert(buf []byte) ([]byte, error) {
	var err error
	s := struct {
		O map[string]any `yaml:",inline"`
	}{}
	err = yaml.Unmarshal(buf, &s)
	if err != nil {
		return nil, err
	}
	buf, err = json.Marshal(s.O)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
