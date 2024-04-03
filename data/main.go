package Data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var artists []Artists
var dates Dates
var Client *http.Client
var Location Locations
var Relations Relation

func GetArtists() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := GetJson(url, &artists)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetDates() {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	err := GetJson(url, &dates)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetJson(url string, target interface{}) error {
	Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetLocations() {
	url := "https://groupietrackers.herokuapp.com/api/location"
	err := GetJson(url, &Location)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetRelations() {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	err := GetJson(url, &Relations)
	if err != nil {
		fmt.Println(err.Error())
	}
}
