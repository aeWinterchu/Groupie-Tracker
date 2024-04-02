package showhomepage

import (
	"Groupie-Tracker/showartistepage"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func ShowHomePage(window fyne.Window) {
	artistsButton := widget.NewButton("Artists", func() {
		showartistepage.ShowArtistsPage()
	})
	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	content := container.NewVBox(
		artistsButton,
		quitButton,
	)
	window.SetContent(content)
}
