package website

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Website struct {
	Icon        string
	Name        string
	Description string
	Url         string
}

var websiteList []*Website

func LoadWebsites() {
	if websiteList == nil {
		websiteList = make([]*Website, 0, 10)
	}
	file, err := os.Open("website/websites.json")
	if err != nil {
		fmt.Println("Failed to read websites.json: ", err.Error())
		return
	}
	fileBytes, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(fileBytes, &websiteList)
	if err != nil {
		fmt.Println("Failed to unmarshall websites.json: ", err.Error())
		return
	}

}
func GetWebsiteList() []*Website {
	if websiteList == nil {
		LoadWebsites()
	}
	return websiteList
}
