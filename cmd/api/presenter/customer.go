package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Customer struct {
	ID        entities.ID `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Address   string      `json:"address"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Latitude  float32     `json:"latitude"`
	Longitude float32     `json:"longitude"`
}
