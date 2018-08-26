package website

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
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
	user, err := user.Current()
	if err != nil {
		log.Fatal("Failed to get current user: ", err.Error())
	}
	file, err := os.Open(user.HomeDir + "/.config/goStudy/websites.json")
	if err != nil {
		log.Fatal("Failed to read websites.json: ", err.Error())
		return
	}
	fileBytes, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(fileBytes, &websiteList)
	if err != nil {
		log.Fatal("Failed to unmarshall websites.json: ", err.Error())
		return
	}

}
func GetWebsiteList() []*Website {
	if websiteList == nil {
		LoadWebsites()
	}
	return websiteList
}
