package configs

import (
	"encoding/json"
	"github.com/hoffme/ddd-backend/internal/apps"
	"os"

	"github.com/hoffme/ddd-backend/internal/contexts"
)

type Config struct {
	Apps     apps.Config     `json:"apps"`
	Contexts contexts.Config `json:"contexts"`
}

func LoadConfigFromJSON(path string) Config {
	config := Config{}

	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(jsonFile).Decode(&config)
	if err != nil {
		panic(err)
	}

	err = jsonFile.Close()
	if err != nil {
		panic(err)
	}

	return config
}
