package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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
		var artistButtons []*widget.Button
		for _, artist := range artists {
			artist := artist // Create a new variable to avoid closure-related issues
			artistButton := widget.NewButton(artist.Name, func() {
				// Handle button click action, for example, open artist's link
				if artist.Link != "" {
					url, err := url.Parse(artist.Link)
					if err != nil {
						fmt.Println("Error parsing URL:", err)
						return
					}
					if err := fyne.CurrentApp().OpenURL(url); err != nil {
						fmt.Println("Error opening URL:", err)
					}
				}
			})
			artistButtons = append(artistButtons, artistButton)
		}

		// Create back button
		var artistObjects []fyne.CanvasObject
for _, artistButton := range artistButtons {
    artistObjects = append(artistObjects, artistButton)
}

// Create back button
backButton := widget.NewButton("Back", func() {
    showHomePage()
})
artistObjects = append(artistObjects, backButton)

scrollableContent := container.NewVBox(
    artistObjects...,
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
