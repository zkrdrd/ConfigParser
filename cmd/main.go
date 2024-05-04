package main

import (
	"fmt"
	"yamljsonread"
)

type HostParser struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func main() {
	var cfg = &HostParser{}
	if err := yamljsonread.Read("config/config.json", cfg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cfg)
	}

}
