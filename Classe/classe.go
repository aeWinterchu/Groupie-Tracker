package Classe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type APIResponse struct {
	Artists   []Artist    `json:"artists"`
	Locations []Locations `json:"locations"`
	Dates     []Dates     `json:"dates"`
	Relation  []Relation  `json:"relation"`
}

type APIResponseDates struct {
	Dates []string `json:"dates"`
}

type APIResponseLocation struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Artist struct {
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int64       `json:"id"`
	DatesLocations []Locations `json:"datesLocations"`
}

func Api_artists() {
	var response []Artist

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range response {
		fmt.Printf("test %d: %s, %s, %d, %s, %s, %s\n", i+1, p.Name, p.Members, p.CreationDate, p.FirstAlbum, p.ConcertDates, p.Image)
	}
}

func Api_dates() {
	var response4 APIResponseDates

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response4)
	if err != nil {
		log.Fatal(err)
	}

	for i, date := range response4.Dates {
		fmt.Printf("test %d: %s\n", i+1, date)
	}
}
func Api_location() {
	var response3 []APIResponseLocation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response3)
	if err != nil {
		var singleResponse APIResponseLocation
		err = json.Unmarshal(body, &singleResponse)
		if err != nil {
			log.Fatal(err)
		}
		response3 = append(response3, singleResponse)
	}

	for i, loc := range response3 {
		fmt.Printf("test %d: %s, %s\n", i+1, loc.Locations, loc.Dates)
	}
}

func Api_Relation() {
	var response2 []Relation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response2)
	if err != nil {
		var singleResponse Relation
		err = json.Unmarshal(body, &singleResponse)
		if err != nil {
			log.Fatal(err)
		}
		response2 = append(response2, singleResponse)
	}

	for i, p := range response2 {
		fmt.Printf("test %d:%v\n", i+1, p.DatesLocations)
	}
}

func newFunction(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

//func pour avoir tout les info de l'artist tout ensemble

func GetArtistInfo(artistName string) {
	// Appel de l'API pour récupérer les données des artistes
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Décode les données JSON de la réponse
	var response APIResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	// Recherche de l'artiste spécifié dans la liste des artistes
	var foundArtist *Artist
	for _, artist := range response.Artists {
		if strings.EqualFold(artist.Name, artistName) {
			foundArtist = &artist
			break
		}
	}

	// Vérifie si l'artiste a été trouvé
	if foundArtist != nil {
		// Affichage des informations sur l'artiste trouvé
		fmt.Printf("Nom: %s\n", foundArtist.Name)
		fmt.Printf("Image: %s\n", foundArtist.Image)
		fmt.Printf("Membres: %s\n", strings.Join(foundArtist.Members, ", "))
		fmt.Printf("Date de création: %d\n", foundArtist.CreationDate)
		fmt.Printf("Premier album: %s\n", foundArtist.FirstAlbum)
		fmt.Printf("Lieux de concerts: %s\n", foundArtist.Locations)
		fmt.Printf("Dates de concerts: %s\n", foundArtist.ConcertDates)
		fmt.Printf("Relations: %s\n", foundArtist.Relations)
	} else {
		fmt.Printf("Artiste '%s' non trouvé.\n", artistName)
	}
}

// /// les filtre
// filtre par le date de la criation de groupe ou artiste
func FilterDate(artists []Artist, fromDate time.Time, toDate time.Time) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		creationDate := time.Unix(artist.CreationDate, 0)
		if creationDate.After(fromDate) && creationDate.Before(toDate) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// filtre par first album
func FilterAlbum(artists []Artist, fromDate time.Time, toDate time.Time) []Artist {
	
		// Define a custom type to hold artist and their first album date
		type ArtistWithFirstAlbumDate struct {
			Artist
			FirstAlbumDate time.Time
		}
	
		// Create a slice to hold artists with their first album date
		var artistsWithFirstAlbumDate []ArtistWithFirstAlbumDate
	
		// Iterate over artists and parse first album date
		for _, artist := range artists {
			var firstAlbumDate time.Time
			// Try different date layouts
			dateFormats := []string{"1000-01-01", "01/01/2024", "1000-01-01T15:04:05Z07:00", "2024-01-01T15:04:05Z"}
			for _, layout := range dateFormats {
				parsedDate, err := time.Parse(layout, artist.FirstAlbum)
				if err == nil {
					firstAlbumDate = parsedDate
					break
				}
			}
			// Add artist and first album date to the slice
			artistsWithFirstAlbumDate = append(artistsWithFirstAlbumDate, ArtistWithFirstAlbumDate{
				Artist:         artist,
				FirstAlbumDate: firstAlbumDate,
			})
		}
	
		// Filter artists based on the specified date range
		var filteredArtists []ArtistWithFirstAlbumDate
		for _, artist := range artistsWithFirstAlbumDate {
			if artist.FirstAlbumDate.After(fromDate) && artist.FirstAlbumDate.Before(toDate) {
				filteredArtists = append(filteredArtists, artist)
			}
		}
	
		// Sort artists by first album date in ascending order
		sort.SliceStable(filteredArtists, func(i, j int) bool {
			return filteredArtists[i].FirstAlbumDate.Before(filteredArtists[j].FirstAlbumDate)
		})
	
		// Extract and return the sorted artists without first album date
		var sortedArtists []Artist
		for _, artist := range filteredArtists {
			sortedArtists = append(sortedArtists, artist.Artist)
		}
	
		return sortedArtists
	}

// filtre pour le concert par localition

