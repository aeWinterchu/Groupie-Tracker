package function

import (
	Data "Main/main.go/data"
	"fmt"
)

var Location Data.Locations

// Retrieve information from the locations and store them in Location
func GetLocations() {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	err := GetJson(url, &Location)
	if err != nil {
		fmt.Println(err.Error())
	}
}
