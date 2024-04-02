package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	// Afficher la page d'accueil
	showHomePage(myApp, myWindow)

	myWindow.ShowAndRun()
}

func showArtistsPage(myApp fyne.App, myWindow fyne.Window) {
	var artists []Artist

	// Récupérer les données des artistes depuis l'API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur lors de la récupération des données, code de statut:", resp.StatusCode)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		fmt.Println("Erreur lors du décodage de la réponse:", err)
		return
	}

	// Créer les boutons d'artiste
	var artistButtons []*widget.Button
	for _, artist := range artists {
		artist := artist
		artistButton := widget.NewButton(artist.Name, func() {
			displayArtistDetails(artist, myApp)
		})
		artistButtons = append(artistButtons, artistButton)
	}

	// Créer le bouton de retour
	backButton := widget.NewButton("Retour", func() {
		showHomePage(myApp, myWindow)
	})

	// Créer la barre de recherche
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher artiste par nom")

	searchButton := widget.NewButton("Rechercher", func() {
		searchTerm := searchEntry.Text
		log.Println("Contenu:", searchTerm)
		searchArtists(artists, searchTerm, myApp)
		searchEntry.SetText("")
	})

	// Contenu de la page des artistes
	content := container.NewVBox(
		searchEntry,
		searchButton,
	)
	for _, artistButton := range artistButtons {
		content.Add(artistButton)
	}
	content.Add(backButton)

	// Rendre le contenu scrollable
	scrollContainer := container.NewScroll(content)
	myWindow.SetContent(scrollContainer)
}

func showHomePage(myApp fyne.App, myWindow fyne.Window) {
	// Créer les boutons pour afficher la liste des artistes et quitter l'application
	artistsButton := widget.NewButton("Artistes", func() {
		showArtistsPage(myApp, myWindow)
	})

	quitButton := widget.NewButton("Quitter", func() {
		myApp.Quit()
	})

	// Contenu de la page d'accueil
	content := container.NewVBox(
		artistsButton,
		quitButton,
	)

	myWindow.SetContent(content)
}

func displayArtistDetails(artist Artist, myApp fyne.App) {
	newWindow := myApp.NewWindow("Détails de l'artiste")
	newWindow.Resize(fyne.NewSize(500, 650))

	artistInfo := widget.NewLabel(fmt.Sprintf("Artiste: %s\nMembres: %s\nVille d'origine: %s\nLien: %s\n", artist.Name, strings.Join(artist.Members, ", "), artist.OriginCity, artist.Link))

	newWindow.SetContent(container.NewScroll(artistInfo))

	newWindow.Show()
}

func searchArtists(artists []Artist, name string, myApp fyne.App) {
	search := strings.ToLower(name)

	var matchingArtists []Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), search) {
			matchingArtists = append(matchingArtists, artist)
		}
	}

	newWindow := myApp.NewWindow("Résultats de la recherche")
	newWindow.Resize(fyne.NewSize(500, 650))

	artistInfo := widget.NewLabel("")

	for _, artist := range matchingArtists {
		artistInfo.SetText(artistInfo.Text + fmt.Sprintf("Artiste: %s\nMembres: %s\nVille d'origine: %s\nLien: %s\n\n", artist.Name, strings.Join(artist.Members, ", "), artist.OriginCity, artist.Link))
	}

	newWindow.SetContent(container.NewScroll(artistInfo))

	newWindow.Show()
}
