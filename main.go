package main

import (
	"time"

	"github.com/joaopedrosgs/PaginaInicialProvas/calendar"

	"github.com/kataras/iris"
)

func main() {
	calendar.Pool()
	go doEvery(time.Hour, calendar.Pool)
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.StaticWeb("/static", "./static")
	app.Get("/", indexHandler)
	app.Run(iris.Addr(":8080"))
}
