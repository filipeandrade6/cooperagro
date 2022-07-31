package domain

import "time"

// type Venue struct {
// 	ID       string
// 	Name     string
// 	Location Location
// }

// type Location struct {
// 	Address          string          `json:"address"`
// 	CrossStreet      string          `json:"crossStreet"`
// 	Latitude         float64         `json:"lat"`
// 	Longitude        float64         `json:"lng"`
// 	LabeledLatLngs   []LabeledCoords `json:"labeledLatLngs"`
// 	Distance         int             `json:"distance"`
// 	PostalCode       string          `json:"postalCode"`
// 	CountryCode      string          `json:"cc"`
// 	City             string          `json:"city"`
// 	State            string          `json:"state"`
// 	Country          string          `json:"country"`
// 	FormattedAddress []string        `json:"formattedAddress"`
// }

// type LabeledCoords struct {
// 	Label string  `json:"label"`
// 	Lat   float64 `json:"lat"`
// 	Lng   float64 `json:"lng"`
// }

type User struct {
	ID        int    `ksql:"id"`
	FirstName string `ksql:"first_name"`
	LastName  string `ksql:"last_name"`
	Email     string `ksql:"email"`
	Phone     string `ksql:"phone"`

	CreatedAt *time.Time `ksql:"created_at"`
	UpdatedAt *time.Time `ksql:"updated_at"`
}

type Product struct {
	ID   int    `ksql:"id"`
	Name string `ksql:"name"`

	CreatedAt *time.Time `ksql:"created_at"`
	UpdatedAt *time.Time `ksql:"updated_at"`
}
