package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type Artist struct {
	ID         int      `json:"id"`
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
	var artists []Artist

	showArtistsPage := func() {
		// Make HTTP request to fetch data from the artists API
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error fetching data, status code:", resp.StatusCode)
			return
		}

		// Decode the response JSON
		if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		// Display fetched data
		var artistButtons []*widget.Button
		for _, artist := range artists {
			artist := artist
			artistButton := widget.NewButton(artist.Name, func() {
				// Handle button click action to display artist details
				displayArtistDetails(artist, myApp)
			})
			artistButtons = append(artistButtons, artistButton)
		}

		// Create back button
		backButton := widget.NewButton("Back", func() {
			showHomePage()
		})

		// Create search bar
		searchEntry := widget.NewEntry()
		searchEntry.SetPlaceHolder("Search artist by name")

		searchButton := widget.NewButton("Search", func() {
			// Call search function here with the search term
			searchTerm := searchEntry.Text
			searchArtists(artists, searchTerm, myApp)
		})

		// Create content container with artist buttons and search bar
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
		myWindow.SetContent(scrollContainer)
	}

	showHomePage = func() {
		// Créez un bouton pour afficher la liste des artistes
		artistsButton := widget.NewButton("Artists", func() {
			showArtistsPage()
		})

		// Créez un bouton pour quitter l'application
		quitButton := widget.NewButton("Quit", func() {
			myApp.Quit()
		})

		// Ajoutez les widgets à la boîte de contenu
		content := container.NewVBox(
			artistsButton,
			quitButton,
		)

		myWindow.SetContent(content)
	}

	showHomePage() // Appelez showHomePage pour afficher la page d'accueil

	myWindow.ShowAndRun()
}

// Func pour afficher les détails de l'artiste
func displayArtistDetails(artist Artist, myApp fyne.App) {
	// Créer une nouvelle fenêtre pour afficher les informations de l'artiste
	newWindow := myApp.NewWindow("Artist Details")
	newWindow.Resize(fyne.NewSize(500, 650))

	// Créer un widget pour afficher les informations de l'artiste
	artistInfo := widget.NewLabel(fmt.Sprintf("Artist: %s\nMembers: %s\nOrigin City: %s\nLink: %s\n", artist.Name, strings.Join(artist.Members, ", "), artist.OriginCity, artist.Link))

	// Ajouter le widget à la fenêtre
	newWindow.SetContent(container.NewScroll(artistInfo))

	// Afficher la nouvelle fenêtre
	newWindow.Show()
}

// Func pour rechercher les artistes par nom
func searchArtists(artists []Artist, name string, myApp fyne.App) {
	// Convertir le nom saisi en minuscules pour une recherche insensible à la casse
	search := strings.ToLower(name)

	// Boucler à travers les artistes pour trouver ceux qui correspondent au terme de recherche
	var matchingArtists []Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), search) {
			matchingArtists = append(matchingArtists, artist)
		}
	}

	// Créer une nouvelle fenêtre pour afficher les résultats de la recherche
	newWindow := myApp.NewWindow("Search Results")
	newWindow.Resize(fyne.NewSize(500, 650))

	// Créer un widget pour afficher les informations des artistes correspondants
	artistInfo := widget.NewLabel("")

	for _, artist := range matchingArtists {
		artistInfo.SetText(artistInfo.Text + fmt.Sprintf("Artist: %s\nMembers: %s\nOrigin City: %s\nLink: %s\n\n", artist.Name, strings.Join(artist.Members, ", "), artist.OriginCity, artist.Link))
	}

	// Ajouter le widget à la fenêtre
	newWindow.SetContent(container.NewScroll(artistInfo))

	// Afficher la nouvelle fenêtre
	newWindow.Show()
}
