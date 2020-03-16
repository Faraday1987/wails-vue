package main

import (
	"wailstest2/pkg/sys"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	stats := &sys.Stats{}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  512,
		Height: 512,
		Title:  "CPU Usage",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(stats)
	app.Run()
}
