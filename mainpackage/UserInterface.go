package mainpackage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

// ShowHomePage displays the home page of the application
func ShowHomePage(myApp fyne.App) {
	// Create an Entry to input search text
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search artist by name")

	// Create a button to trigger the search
	searchButton := widget.NewButton("Search", func() {
		// Call the searchBarre function with the list of artists and the entered text as arguments
		searchBarre(artists, searchEntry.Text)
	})

	// Create a button to display the list of artists
	artistsButton := widget.NewButton("Artists", func() {
		ShowArtistsPage(myApp)
	})

	// Content container
	content := container.NewVBox(
		widget.NewLabel("Hello!"),
		searchEntry,
		searchButton,
		artistsButton,
	)

	myApp.SetContent(content)
}

// ShowArtistsPage displays the artists page
func ShowArtistsPage(myApp fyne.App) {
	// Make an HTTP request to fetch data from the artists API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Display fetched data
	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	var artistButtons []*widget.Button
	for _, artist := range artists {
		artist := artist // Avoid closure-related issues
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

	// Back button
	backButton := widget.NewButton("Back", func() {
		ShowHomePage(myApp)
	})

	// Search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search artist by name")

	searchButton := widget.NewButton("Search", func() {
		// Call your search function here with the search term
		searchTerm := searchEntry.Text
		searchForArtist(searchTerm)
	})

	// Content container with artist buttons and search bar
	content := container.NewVBox(
		searchEntry,
		searchButton,
	)
	for _, artistButton := range artistButtons {
		content.Add(artistButton)
	}
	content.Add(backButton)

	// Make the content scrollable
	scrollContainer := container.NewScroll(content)
	myApp.SetContent(scrollContainer)
}

// ShowArtistDetails displays details of a specific artist
func ShowArtistDetails(myApp fyne.App, artist Artist) {
	nameLabel := widget.NewLabel("Name: " + artist.Name)
	membersLabel := widget.NewLabel("Members: " + strings.Join(artist.Members, ", "))
	originLabel := widget.NewLabel("Origin city: " + artist.OriginCity)
	linkLabel := widget.NewLabel("Link: " + artist.Link)

	// Back button
	backButton := widget.NewButton("Back", func() {
		ShowArtistsPage(myApp)
	})

	// Content container
	content := container.NewVBox(
		nameLabel,
		membersLabel,
		originLabel,
		linkLabel,
		backButton,
	)
	myApp.SetContent(content)
}
