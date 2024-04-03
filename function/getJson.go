package function

import (
	"encoding/json"
	"net/http"
	"time"
)

var Client *http.Client

// Create a function to connect to the API and take the URL of the page and the variable in which to store the retrieved API information as parameters
func GetJson(url string, target interface{}) error {
	Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
