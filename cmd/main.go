package main

import (
	"fmt"
	"yamljsonread"
)

func main() {
	var c yamljsonread.ConfigReader = &yamljsonread.HostParser{}
	r := c.Read("configs/config.json")
	fmt.Println(r)
}
