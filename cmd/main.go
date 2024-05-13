package main

import (
	"fmt"

	parser "github.com/zkrdrd/config-parser.git"
)

type HostParser struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func main() {
	var cfg = &HostParser{}
	if err := parser.Read("configs/config.json", cfg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cfg)
	}

}
