package parser

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func Read(Filename string, cfg any) error {
	file, err := os.ReadFile(Filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, cfg); err == nil {
		return nil
	}

	if err := yaml.Unmarshal(file, cfg); err == nil {
		return nil
	}

	return err
}
