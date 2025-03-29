package byebyelinktracking

import (
	"encoding/json"
	"errors"
	"strings"

	"sigs.k8s.io/yaml"
)

type Config struct {
	Entries []configEntry `json:"entries"`
}

func parseConfig(data []byte, ext string) (Config, error) {
	c := Config{}

	var err error

	switch strings.ToLower(ext) {
	case ".json":
		err = json.Unmarshal(data, &c)
	case ".yaml":
		err = yaml.Unmarshal(data, &c)
	default:
		err = errors.New("no supported config type used")
	}

	return c, err
}
