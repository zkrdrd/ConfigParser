package parser_test

import (
	"fmt"
	"testing"
	parser "yamljsonread"
)

type HostParser struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func TestRead(t *testing.T) {
	t.Parallel()
	t.Run(`test pase file`, func(t *testing.T) {
		testTable := []struct {
			Values      string
			ErrorResult string
			Result      *HostParser
		}{
			{
				Values: "configs/config.json",
				Result: &HostParser{
					Host: "1.1.1.1",
					Port: 123,
				},
			},
			{
				Values: "configs/config.yaml",
				Result: &HostParser{
					Host: "2.2.2.2",
					Port: 456,
				},
			},
			{
				Values: "configs/config",
				Result: &HostParser{
					Host: "1.1.1.1",
					Port: 123,
				},
			},
			{
				Values: "configs/config2",
				Result: &HostParser{
					Host: "2.2.2.2",
					Port: 456,
				},
			},
			{
				Values:      "configs/noExistFile",
				ErrorResult: "open configs/noExistFile: no such file or directory",
			},
			{
				Values:      "configs/empty",
				ErrorResult: "file is empty",
			},
			{
				Values:      "configs/noFormat.txt",
				ErrorResult: "parser not found",
			},
		}

		var cfg = &HostParser{}
		for _, expect := range testTable {
			if err := parser.Read(expect.Values, cfg); err != nil {
				if expect.ErrorResult != error.Error(err) {
					t.Error(fmt.Errorf(`another error want %v != %v`, expect.ErrorResult, err))
				}
				continue
			} else {
				if expect.Result.Host != cfg.Host {
					t.Error(fmt.Errorf(`host is not current want %v != %v`, expect.Result.Host, cfg.Host))
				}
				if expect.Result.Port != cfg.Port {
					t.Error(fmt.Errorf(`port is not current want %v != %v`, expect.Result.Host, cfg.Port))
				}
			}
		}
	})
}
