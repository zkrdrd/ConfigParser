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
			HostResult  string
			PortResult  int
		}{
			{
				Values:     "configs/config.json",
				HostResult: "1.1.1.1",
				PortResult: 123,
			},
			{
				Values:     "configs/config.yaml",
				HostResult: "2.2.2.2",
				PortResult: 456,
			},
			{
				Values:     "configs/config",
				HostResult: "1.1.1.1",
				PortResult: 123,
			},
			{
				Values:     "configs/config2",
				HostResult: "2.2.2.2",
				PortResult: 456,
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
				resultParse := parser.Read(expect.Values, cfg)
				if resultParse != nil {
					if expect.ErrorResult != fmt.Sprintf("%v", resultParse) {
						t.Error(fmt.Errorf(`result %v != %v`, expect.HostResult, resultParse))
					}
				} else {
					if expect.HostResult != cfg.Host {
						t.Error(fmt.Errorf(`result %v != %v`, expect.HostResult, cfg.Host))
					}
					if expect.PortResult != cfg.Port {
						t.Error(fmt.Errorf(`result %v != %v`, expect.PortResult, cfg.Port))
					}
				}
			}()
		}
		wg.Wait()
	})

}
