package jsonrd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

func JsonDecode() {
	jsonFile, err := os.Open("usrs.json")
	if err != nil {
		fmt.Printf("error : %s", err)
	}
	defer jsonFile.Close()
	var usrs Usrs
	byt3, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byt3, &usrs)
	for index := 0; index < len(usrs.Usrs); index++ {
		fmt.Println("User Type: " + usrs.Usrs[index].Type)
		fmt.Println("User Age: " + strconv.Itoa(usrs.Usrs[index].Age))
		fmt.Println("User Name: " + usrs.Usrs[index].Name)
		fmt.Println("Facebook Url: " + usrs.Usrs[index].Social.Facebook)
		fmt.Println("Twitter Url:" + usrs.Usrs[index].Social.Twitter)
	}
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
