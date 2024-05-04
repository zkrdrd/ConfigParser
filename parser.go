package parser

import (
	"encoding/json"
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func Read(Filename string, cfg any) error {
	fileData, err := os.ReadFile(Filename)
	if err != nil {
		return err
	}
	if len(fileData) == 0 {
		return errors.New("file is empty")
	}

	if err := json.Unmarshal(fileData, cfg); err == nil {
		return nil
	}

	if err := yaml.Unmarshal(fileData, cfg); err == nil {
		return nil
	}

	return errors.New("parser not found")
}
