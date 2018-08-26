package main

import (
	"github.com/joaopedrosgs/GoStudy/calendar"
	"github.com/joaopedrosgs/GoStudy/website"
	"github.com/kataras/iris"
)

func indexHandler(ctx iris.Context) {
	websites := website.GetWebsiteList()
	events := calendar.GetEventList()
	ctx.ViewData("websiteList", websites)
	ctx.ViewData("eventList", events)
	ctx.View("index.html")
}
