package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kataras/iris"
)

/*
if len(events.Items) == 0 {
	fmt.Println("No upcoming events found.")
} else {
	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		timeDate, _ := time.Parse(time.RFC3339, date)
		if strings.Contains(item.Summary, "Prova") {
			fmt.Printf("%v (%s)\n", item.Summary, humanize.Time(timeDate))
		}
	}
}*/
func main() {
	UpdateEvents()
	go doEvery(time.Hour, UpdateEvents)
	app := iris.New()

	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)
	app.StaticWeb("/static", "./static")
	app.Get("/", indexHandler)
	app.Build()
	app.Run(iris.Addr(":8080"))
}
func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func getSites() []string {
	var listaSites map[string][]string
	jsonFile, err := os.Open("sites.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &listaSites)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return listaSites["sites"]

}

func indexHandler(ctx iris.Context) {
	// Bind: {{.message}} with "Hello world!"
	ctx.ViewData("sites", getSites())
	ctx.ViewData("eventos", eventos)
	// Render template file: ./views/hello.html
	ctx.View("index.html")
}
