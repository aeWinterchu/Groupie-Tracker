package showhomepage

import (
    "fyne.io/fyne"
    "fyne.io/fyne/container"
    "fyne.io/fyne/widget"
    "Groupie-Tracker/showartistepage"
)

func ShowHomePage(window fyne.Window) {
    artistsButton := widget.NewButton("Artists", func() {
        showartistepage.ShowArtistsPage()
    })
    quitButton := widget.NewButton("Quit", func() {
        myApp.Quit()
    })

    content := container.NewVBox(
        widget.NewLabel("Hello!"),
        artistsButton,
        quitButton,
    )
    window.SetContent(content)
}