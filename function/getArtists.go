package function

import (
	Data "Main/main.go/data"
	"fmt"
)

var Artists []Data.Artists

// Retrieve information from the artists page and store it in the Artists variable
func GetArtists() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := GetJson(url, &Artists)
	if err != nil {
		fmt.Println(err.Error())
	}
}
