package main

import (
	// "image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	manGo:= app.New()
	icon, _ := fyne.LoadResourceFromPath("icons\\mangologo.png")
	myWindow := manGo.NewWindow("ManGo")
	myWindow.Resize(fyne.NewSize(500, 500))
	//myWindow.SetFixedSize(true)
	background := canvas.NewImageFromFile("backgrounds\\mangoback.png")
	Welcome := widget.NewLabel("Welcome to ManGO")
	centered_welcome := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), Welcome, layout.NewSpacer())
	descargar := widget.NewButton("Descargar", func() {
	})
	descargar.Resize(fyne.NewSize(150,30))
	buttons := container.New(
		layout.NewHBoxLayout(), 
		layout.NewSpacer(), 
		widget.NewButton("Descargar", func() {
		}),
		widget.NewButton("Cerrar", func ()  {
			myWindow.Close()
		}),descargar,
		layout.NewSpacer())
	backgroundc := container.New(layout.NewMaxLayout(),background, centered_welcome, buttons)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), backgroundc,centered_welcome, buttons))
	myWindow.SetContent(backgroundc)
	myWindow.SetIcon(icon)
	myWindow.ShowAndRun()
}