package main

import (
	"gtkgo/client/views"

	"fyne.io/fyne/v2/app"
)

func main() {

	app := app.New()
	app.NewWindow("Fyne App")

	client := views.NewRegisterWindow(app)
	client.BuildAndShow()

	app.Run()
}
