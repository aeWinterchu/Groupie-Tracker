package function

import (
	Data "Main/main.go/data"
	"fmt"
)

var Dates Data.Dates

// Retrieve information from the dates page of the API and store it in Dates
func GetDates() {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	err := GetJson(url, &Dates)
	if err != nil {
		fmt.Println(err.Error())
	}
}
