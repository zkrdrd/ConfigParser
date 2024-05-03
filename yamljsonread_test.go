package yamljsonread_test

import (
	"fmt"
	"sync"
	"testing"
	"yamljsonread"
)

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
			{
				Values: "configs/confignull",
				Result: "open configs/confignull: no such file or directory",
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наименьшего числа
				var result yamljsonread.ConfigReader = &yamljsonread.HostParser{}
				c := result.Read(expect.Values)
				if expect.Result != c {
					t.Error(fmt.Errorf(`result %v != %v`, expect.Result, c))
				}
			}()
		}

		wg.Wait()
	})

}
