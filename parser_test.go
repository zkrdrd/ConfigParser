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
				Values:      "configs/notExistFile",
				ErrorResult: "open configs/notExistFile: no such file or directory",
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
				// Проверяем ошибку
				if expect.ErrorResult != err.Error() {
					t.Error(fmt.Errorf(`ожидалась другая ошибка %v != %v`, expect.ErrorResult, err))
				}
				continue
			}

			if expect.Result.Host != cfg.Host {
				t.Error(fmt.Errorf(`host is not current want: %s, have: %s`, expect.Result.Host, cfg.Host))
			}

			if expect.Result.Port != cfg.Port {
				t.Error(fmt.Errorf(`host is not current want: %d, have: %d`, expect.Result.Port, cfg.Port))
			}
		}
	})

}
