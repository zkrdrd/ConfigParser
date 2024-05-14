package main

import (
	"fmt"

	parser "github.com/zkrdrd/ConfigParser"
)

type HostParser struct {
	Host string `json:"host" yaml:"host" xml:"Host" toml:"host"`
	Port int    `json:"port" yaml:"port" xml:"Port" toml:"port"`
}

func main() {
	var cfg = &HostParser{}
	if err := parser.Read("configs/config.toml", cfg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cfg)
	}

}
