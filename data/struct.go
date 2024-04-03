package Data

// Create a structure adapted to the API to retrieve and use information
type Artists struct {
	Id           int      `json:"id"`           // Artist ID
	Image        string   `json:"image"`        // Artist image URL
	Name         string   `json:"name"`         // Artist name
	Members      []string `json:"members"`      // Members of the artist group
	CreationDate int      `json:"creationDate"` // Date of artist creation
	FirstAlbum   string   `json:"firstAlbum"`   // First album release date
	Location     string   `json:"locations"`    // Artist's location
	ConcertDates string   `json:"concertDates"` // Dates of concerts
	Relations    string   `json:"relations"`    // Relations with other artists
}

type DatesInfo struct {
	Id   int      `json:"id"`    // ID of the concert date
	Date []string `json:"dates"` // List of dates
}

type Dates struct {
	Index []DatesInfo `json:"index"` // Index of DatesInfo
}

type Locations struct {
	Index []LocationsInfo `json:"index"` // Index of LocationsInfo
}

type LocationsInfo struct {
	Id       int      `json:"id"`        // ID of the location
	Location []string `json:"locations"` // List of locations
	Date     string   `json:"dates"`     // Date associated with the location
}

type Relation struct {
	Index []RelationInfo `json:"index"` // Index of RelationInfo
}

type RelationInfo struct {
	Id              int                 `json:"id"`             // ID of the relation
	Dates_Locations map[string][]string `json:"datesLocations"` // Mapping of dates to locations
}
