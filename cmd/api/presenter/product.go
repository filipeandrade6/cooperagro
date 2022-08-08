package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Product struct {
	ID            entities.ID `json:"id"`
	Name          string      `json:"name"`
	BaseProductID entities.ID `json:"base_product_id"`
}

type CreateProduct struct {
	Name          string      `json:"name" binding:"required"`
	BaseProductID entities.ID `json:"base_product_id" binding:"required"`
}

type UpdateProduct struct {
	Name          string      `json:"name" binding:"required"`
	BaseProductID entities.ID `json:"base_product_id" binding:"required"`
}
