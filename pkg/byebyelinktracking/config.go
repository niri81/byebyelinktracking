package byebyelinktracking

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Config struct {
	Entries []configEntry `json:"entries"`
}

func parseConfig(data []byte) Config {
	c := Config{}

	err := json.Unmarshal(data, &c)
	if err != nil {
		slog.Error("could not parse config", "err", err.Error())
		os.Exit(1)
	}

	return c
}
