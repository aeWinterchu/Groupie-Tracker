package showartistepage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"Groupie-Tracker/showHomePage"
)

type Artist struct {
	Name       string   `json:"name"`
	Members    []string `json:"members"`
	OriginCity string   `json:"origin_city"`
	Link       string   `json:"link"`
}

// Fonction pour afficher la page des artistes
func ShowArtistsPage(myWindow fyne.Window) {
	var artists []Artist
	// Faire une requête HTTP pour récupérer les données depuis l'API des artistes
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données :", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		fmt.Println("Erreur lors du décodage de la réponse :", err)
		return
	}

	// Afficher les données récupérées
	var artistButtons []*widget.Button
	for _, artist := range artists {
		artist := artist // Créer une nouvelle variable pour éviter les problèmes de fermeture
		artistButton := widget.NewButton(artist.Name, func() {
			// Gérer l'action du clic sur le bouton, par exemple, ouvrir le lien de l'artiste
			if artist.Link != "" {
				url, err := url.Parse(artist.Link)
				if err != nil {
					fmt.Println("Erreur lors de l'analyse de l'URL :", err)
					return
				}
				if err := fyne.CurrentApp().OpenURL(url); err != nil {
					fmt.Println("Erreur lors de l'ouverture de l'URL :", err)
				}
			}
		})
		artistButtons = append(artistButtons, artistButton)
	}

	// Créer le bouton de retour
	var artistObjects []fyne.CanvasObject
	for _, artistButton := range artistButtons {
		artistObjects = append(artistObjects, artistButton)
	}

	// Créer le bouton de retour
	backButton := widget.NewButton("Retour", func() {
		showhomepageartiste.ShowHomePage(myWindow) // Appeler showHomePage avec la référence de la fenêtre
	})
	artistObjects = append(artistObjects, backButton)

	// Ajouter la barre de recherche uniquement sur la page des artistes
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste par nom")

	searchButton := widget.NewButton("Rechercher", func() {
		searchBar(artists, searchEntry.Text)
	})

	artistObjects = append(artistObjects, searchEntry, searchButton)

	scrollableContent := container.NewVBox(
		artistObjects...,
	)

	  // Rendre le contenu défilable
	scrollContainer := container.NewScroll(scrollableContent)
	myWindow.SetContent(scrollContainer)
}
