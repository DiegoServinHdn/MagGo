package main

import (
	// "image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	manGo := app.New()
	icon, _ := fyne.LoadResourceFromPath("icons\\mangologo.png")
	myWindow := manGo.NewWindow("ManGo")
	myWindow.Resize(fyne.NewSize(500, 500))
	Welcome := widget.NewLabel("Welcome to ManGO")
	centered_welcome := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), Welcome, layout.NewSpacer())
	buttons := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		widget.NewButton("Descargar", func() {
			Welcome.SetText("hi")
		}),
		widget.NewButton("Cerrar", func() {
			myWindow.Close()
		}),
		layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), centered_welcome, buttons))

	myWindow.SetIcon(icon)
	myWindow.ShowAndRun()
}
