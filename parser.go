package parser

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

type parser func([]byte, any) error

var (
	parsers           []parser
	ErrFileIsEmpty    = errors.New(`file is empty`)
	ErrParsernotFound = errors.New(`parser not found`)
)

func init() {
	parsers = append(parsers, json.Unmarshal)
	parsers = append(parsers, yaml.Unmarshal)
}

func Read(Filename string, cfg any) error {
	fileData, err := os.ReadFile(Filename)
	if err != nil {
		return fs.ErrNotExist
	}

	if len(fileData) == 0 {
		return ErrFileIsEmpty
	}

	for _, parserType := range parsers {
		if err := parserType(fileData, cfg); err == nil {
			return nil
		}
	}

	return ErrParsernotFound
}
