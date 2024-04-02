package mainpackage

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"strings"
	"Groupie-Tracker\Classe"
)

type Artist struct {
	Name  string `json:"name"`
	Membres  []string `json:"membres"`
	OriginCity    string    `json:"origin_city"`
	Link          string      `json:"link"`

}


func searchForArtist(searchTerm string){
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

    // Filter artists based on the search term
    var filteredArtists []Artist
    for _, artist := range artists {
        if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchTerm)) {
            filteredArtists = append(filteredArtists, artist)
        }
    }
}


func searchBarre(artists []Artist, name string){
	 // Convertir le nom saisi en minuscules pour une recherche insensible à la casse
	 search := strings.ToLower(name)

	 // Boucler à travers les artistes pour trouver ceux qui correspondent au terme de recherche
	 var correctArtists []Artist
	 
	 for _, artist := range artists {
		 if strings.Contains(strings.ToLower(artist.Name), search) {
			 correctArtists = append(correctArtists, artist)
			 GetArtistInfo(name)
		 }
	 }
 
	 // Faire quelque chose avec les artistes correspondants
	 fmt.Println(correctArtists)
}
