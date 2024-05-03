package yamljsonread

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigReader interface {
	Read(Filename string) string
}

type HostParser struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func (cfg *HostParser) Read(Filename string) string {
	file, err := os.Open(Filename)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	defer file.Close()
	value, err := io.ReadAll(file)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	err = json.Unmarshal(value, &cfg)
	if err != nil {
		err := yaml.Unmarshal(value, &cfg)
		if err != nil {
			return fmt.Sprintf("%v", err)
		} else {
			return fmt.Sprintf("%v", cfg)
		}
	} else {
		return fmt.Sprintf("%v", cfg)
	}
}
