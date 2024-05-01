package csvreader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Csvreader interface {
	Read(Filename string) string
}

var (
	Filename string
)

type HostParser struct {
	HostParser []Host `json:"host"`
}

type Host struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func (hp *HostParser) Read(Filename string) string {
	file, err := os.Open(Filename)
	if err != nil {
		panic("Error open file")
	}
	defer file.Close()
	value, _ := io.ReadAll(file)
	json.Unmarshal(value, &hp)
	//fmt.Println(json.Unmarshal(value, &hp))
	//var result map[string]interface{}
	//json.Unmarshal([]byte(value), &result)
	for i := 0; i < (len(hp.HostParser)); i++ {
		fmt.Printf("ip: %s \n", string(hp.HostParser[i].Ip))
		fmt.Printf("port: %s \n", string(hp.HostParser[i].Port))
	}
	//fmt.Print(result["host"])
	return ""
}
