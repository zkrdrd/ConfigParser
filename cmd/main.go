package main

import (
	"yamljsonread"
)

func main() {
	var c yamljsonread.Csvreader = &yamljsonread.HostParser{}
	c.Read("config.json")
	//fmt.Println(c)
}
