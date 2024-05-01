package main

import (
	"csvreader"
)

func main() {
	var c csvreader.Csvreader = &csvreader.HostParser{}
	c.Read("config.json")
	//fmt.Println(c)
}
