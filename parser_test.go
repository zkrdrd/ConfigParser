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
	t.Run(`test find min number`, func(t *testing.T) {
		testTable := []struct {
			Values string
			Result string
		}{
			{
				Values: "configs/config.json",
				Result: "&{1.1.1.1 123}",
			},
			{
				Values: "configs/config.yaml",
				Result: "&{2.2.2.2 456}",
			},
			{
				Values: "configs/config",
				Result: "&{1.1.1.1 123}",
			},
			{
				Values: "configs/config2",
				Result: "&{2.2.2.2 456}",
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наименьшего числа
				var cfg = &HostParser{}
				c := parser.Read(expect.Values, cfg)

				if expect.Result != fmt.Sprintf("%v", cfg) {
					t.Error(fmt.Errorf(`result %v != %v`, expect.Result, c))
				}
			}()
		}

		wg.Wait()
	})

}
