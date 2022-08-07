package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Inventory struct {
	ID            entities.ID `json:"id"`
	User          entities.ID `json:"user"`
	Product       entities.ID `json:"product"`
	Quantity      int         `json:"quantity"`
	UnitOfMeasure entities.ID `json:"unit_of_measure"`
}

type CreateInventory struct {
	User          entities.ID `json:"user" binding:"required"`
	Product       entities.ID `json:"product" binding:"required"`
	Quantity      int         `json:"quantity" binding:"required"`
	UnitOfMeasure entities.ID `json:"unit_of_measure" binding:"required"`
}

type UpdateInventory struct {
	User          entities.ID `json:"user" binding:"required"`
	Product       entities.ID `json:"product" binding:"required"`
	Quantity      int         `json:"quantity" binding:"required"`
	UnitOfMeasure entities.ID `json:"unit_of_measure" binding:"required"`
}
