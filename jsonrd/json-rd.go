package jsonrd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func JsonReader() {
	jsonFile, err := os.Open("usrs.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var rs map[string]interface{}
	json.Unmarshal([]byte(byteValue), &rs)
	fmt.Println(rs["users"])
}

type Usrs struct {
	Usrs []Usr `json:"users"`
}

type Usr struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
