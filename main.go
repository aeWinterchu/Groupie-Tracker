package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type Artist struct {
	Name       string   `json:"name"`
	Members    []string `json:"members"`
	OriginCity string   `json:"origin_city"`
	Link       string   `json:"link"`
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie-Tracker")
	myWindow.Resize(fyne.NewSize(500, 650))

	var showHomePage func()
	showArtistsPage := func() {
		// Make HTTP request to fetch data from the artists API
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return
		}
		defer resp.Body.Close()

		var artists []Artist
		if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		// Display fetched data
		var artistsContent []fyne.CanvasObject
		for _, artist := range artists {
			artistLabel := widget.NewLabel(fmt.Sprintf("Name: %s", artist.Name))
			artistsContent = append(artistsContent, artistLabel)
		}

		// Create back button
		backButton := widget.NewButton("Back", func() {
			showHomePage()
		})

		artistsContent = append(artistsContent, backButton)

		scrollableContent := container.NewVBox(
			artistsContent...,
		)

		// Make the content scrollable
		scrollContainer := container.NewScroll(scrollableContent)

		myWindow.SetContent(scrollContainer)
	}
	showHomePage = func() {
		// Create buttons
		artistsButton := widget.NewButton("Artists", func() {
			showArtistsPage()
		})
		quitButton := widget.NewButton("Quit", func() {
			myApp.Quit()
		})

		myWindow.SetContent(container.NewVBox(
			widget.NewLabel("Hello!"),
			artistsButton,
			quitButton,
		))
	}
	showHomePage()

	myWindow.ShowAndRun()
}
