package presenter

import "github.com/filipeandrade6/cooperagro/domain/entity"

type Inventory struct {
	ID              entity.ID `json:"id"`
	UserID          entity.ID `json:"user_id"`
	ProductID       entity.ID `json:"product_id"`
	Quantity        int       `json:"quantity"`
	UnitOfMeasureID entity.ID `json:"unit_of_measure_id"`
}

type CreateInventory struct {
	UserID          entity.ID `json:"user_id" binding:"required"`
	ProductID       entity.ID `json:"product_id" binding:"required"`
	Quantity        int       `json:"quantity" binding:"required"`
	UnitOfMeasureID entity.ID `json:"unit_of_measure_id" binding:"required"`
}

type UpdateInventory struct {
	UserID          entity.ID `json:"user_id" binding:"required"`
	ProductID       entity.ID `json:"product_id" binding:"required"`
	Quantity        int       `json:"quantity" binding:"required"`
	UnitOfMeasureID entity.ID `json:"unit_of_measure_id" binding:"required"`
}

type EchoInventory struct {
	UserID          string `json:"user_id"`
	ProductID       string `json:"product_id"`
	Quantity        int    `json:"quantity"`
	UnitOfMeasureID string `json:"unit_of_measure_id"`
}

type EchoCreateInventory struct {
	UserID          string `json:"user_id"`
	ProductID       string `json:"product_id"`
	Quantity        int    `json:"quantity"`
	UnitOfMeasureID string `json:"unit_of_measure_id"`
}

type EchoUpdateInventory struct {
	UserID          string `json:"user_id"`
	ProductID       string `json:"product_id"`
	Quantity        int    `json:"quantity"`
	UnitOfMeasureID string `json:"unit_of_measure_id"`
}

type EchoDeleteInventory struct {
	ID string `json:"id"`
}
