package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	manGO := app.New()
	home:= manGO.NewWindow("Mango")

	download := widget.NewLabel("Hello ManGo")
	home.SetContent(container.NewVBox(
		download,
		widget.NewButton("Descargar", func() {
			download.SetText("Welcome :)")
		}),
	))

	home.ShowAndRun()
}