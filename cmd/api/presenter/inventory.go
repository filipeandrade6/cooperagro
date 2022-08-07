package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Inventory struct {
	ID            entities.ID `json:"id"`
	Customer      entities.ID `json:"customer"`
	Product       entities.ID `json:"product"`
	Quantity      int         `json:"quantity"`
	UnitOfMeasure entities.ID `json:"unit_of_measure"`
}
