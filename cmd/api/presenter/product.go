package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Product struct {
	ID          entities.ID `json:"id"`
	Name        string      `json:"name"`
	BaseProduct entities.ID `json:"base_product"`
}

type CreateProduct struct {
	Name        string      `json:"name" binding:"required"`
	BaseProduct entities.ID `json:"base_product" binding:"required"`
}

type UpdateProduct struct {
	Name        string      `json:"name" binding:"required"`
	BaseProduct entities.ID `json:"base_product" binding:"required"`
}
