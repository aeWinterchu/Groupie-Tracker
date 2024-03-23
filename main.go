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
	"strings"
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
	var artists []Artist
	showHomePage = func() {
        // Créez une Entry pour saisir le texte de recherche
        searchEntry := widget.NewEntry()
        searchEntry.SetPlaceHolder("Search artist by name")

        // Créez un bouton pour déclencher la recherche
        searchButton := widget.NewButton("Search", func() {
            // Appeler la fonction searchBarre avec la liste des artistes et le texte saisi comme arguments
            searchBarre(artists, searchEntry.Text)
        })

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
            widget.NewLabel("Hello!"),
            searchEntry,
            searchButton,
            artistsButton,
            quitButton,
        )

        myWindow.SetContent(content)
    }

    showHomePage() // Appelez showHomePage pour afficher la page d'accueil

    myWindow.ShowAndRun()
}


// Func pour la barre de recherche
func searchBarre(artists []Artist,name string) {
    // Convertir le nom saisi en minuscules pour une recherche insensible à la casse
    search := strings.ToLower(name)

    // Boucler à travers les artistes pour trouver ceux qui correspondent au terme de recherche
    var correctArtists []Artist
    
    for _, artist := range artists {
        if strings.Contains(strings.ToLower(artist.Name), search) {
            correctArtists = append(correctArtists, artist)
        }
    }

    // Faire quelque chose avec les artistes correspondants
    fmt.Println(correctArtists)
}
