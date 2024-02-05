package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie-Tracker")

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Hello!"),
		widget.NewButton("Artists", func() {

		}),
		widget.NewButton("Locations", func() {

		}),
		widget.NewButton("Dates", func() {

		}),
		widget.NewButton("Quit", func() {
			myApp.Quit()
		})))

	myWindow.ShowAndRun()
}
