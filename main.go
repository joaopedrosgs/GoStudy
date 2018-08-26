package main

import (
	"log"
	"os/user"
	"time"

	"github.com/joaopedrosgs/GoStudy/calendar"

	"github.com/kataras/iris"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("failed to get current user: ", err.Error())
	}
	calendar.Pool()
	go doEvery(time.Hour, calendar.Pool)
	app := iris.New()
	tmpl := iris.HTML(user.HomeDir+"/.local/share/go-study/views", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.StaticWeb("/static", user.HomeDir+"/.local/share/go-study/static")
	app.Get("/", indexHandler)
	app.Run(iris.Addr(":8080"))
}
