package parser_test

import (
	"fmt"
	"sync"
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
			Result      HostParser
		}{
			{
				Values: "configs/config.json",
				Result: HostParser{
					Host: "1.1.1.1",
					Port: 123,
				},
			},
			{
				Values: "configs/config.yaml",
				Result: HostParser{
					Host: "2.2.2.2",
					Port: 456,
				},
			},
			{
				Values: "configs/config",
				Result: HostParser{
					Host: "1.1.1.1",
					Port: 123,
				},
			},
			{
				Values: "configs/config2",
				Result: HostParser{
					Host: "2.2.2.2",
					Port: 456,
				},
			},
			{
				Values:      "configs/noFile",
				ErrorResult: "open configs/noFile: no such file or directory",
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

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				var cfg = &HostParser{}
				if resultParse := parser.Read(expect.Values, cfg); resultParse != nil {
					if expect.ErrorResult != error.Error(resultParse) {
						t.Error(fmt.Errorf(`result %v != %v`, expect.ErrorResult, resultParse))
					}
				} else {
					if expect.Result != *cfg {
						t.Error(fmt.Errorf(`result %v != %v`, expect.Result.Host, cfg.Host))
					}
				}
			}()
		}
		wg.Wait()
	})

}
