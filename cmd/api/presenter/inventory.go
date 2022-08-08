package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Inventory struct {
	ID              entities.ID `json:"id"`
	UserID          entities.ID `json:"user_id"`
	ProductID       entities.ID `json:"product_id"`
	Quantity        int         `json:"quantity"`
	UnitOfMeasureID entities.ID `json:"unit_of_measure_id"`
}

type CreateInventory struct {
	UserID          entities.ID `json:"user_id" binding:"required"`
	ProductID       entities.ID `json:"product_id" binding:"required"`
	Quantity        int         `json:"quantity" binding:"required"`
	UnitOfMeasureID entities.ID `json:"unit_of_measure_id" binding:"required"`
}

type UpdateInventory struct {
	UserID          entities.ID `json:"user_id" binding:"required"`
	ProductID       entities.ID `json:"product_id" binding:"required"`
	Quantity        int         `json:"quantity" binding:"required"`
	UnitOfMeasureID entities.ID `json:"unit_of_measure_id" binding:"required"`
}
