package yamljsonread

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
		return json.Unmarshal(file, cfg)
	}

	if err := yaml.Unmarshal(file, cfg); err == nil {
		return yaml.Unmarshal(file, cfg)
	}

	return err
}
