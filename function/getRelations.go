package function

import (
	Data "Main/main.go/data"
	"fmt"
)

var Relations Data.Relation

// Retrieve information from the API of the relation page and store them in the variable Relations of type Data.Relation
func GetRelations() {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	err := GetJson(url, &Relations)
	if err != nil {
		fmt.Println(err.Error())
	}
}
