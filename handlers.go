package main

import (
	"github.com/joaopedrosgs/PaginaInicialProvas/calendar"
	"github.com/joaopedrosgs/PaginaInicialProvas/website"
	"github.com/kataras/iris"
)

func indexHandler(ctx iris.Context) {
	websites := website.GetWebsiteList()
	events := calendar.GetEventList()
	ctx.ViewData("websiteList", websites)
	ctx.ViewData("eventList", events)
	ctx.View("index.html")
}
