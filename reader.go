package yamljsonread

import (
	//"encoding/json"

	"encoding/json"
	"fmt"
	"io"
	"os"
	//"gopkg.in/yaml.v3"
)

type Csvreader interface {
	Read(Filename string)
}
type HostParser struct{}

/*
	type HostParser struct {
		HostParser []Host `json:"host"`
	}

	type Host struct {
		Ip   string `json:"ip"`
		Port string `json:"port"`
	}
*/
func (hp *HostParser) Read(Filename string) {
	file, err := os.Open(Filename)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	value, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}
	//err = json.Unmarshal(value, &hp)
	fmt.Print(json.Unmarshal(value, &hp))
	//err = yaml.Unmarshal(value, &hp)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(value))
	//fmt.Println(json.Unmarshal(value, &hp))
	//var result map[string]interface{}
	//json.Unmarshal([]byte(value), &result)
	/*for i := 0; i < (len(hp.HostParser)); i++ {
		fmt.Printf("ip: %s \n", string(hp.HostParser[i].Ip))
		fmt.Printf("port: %s \n", string(hp.HostParser[i].Port))
	}*/
	//fmt.Print(result["host"])

}
