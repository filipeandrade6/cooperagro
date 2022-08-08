package presenter

import "github.com/filipeandrade6/cooperagro/domain/entity"

type Product struct {
	ID            entity.ID `json:"id"`
	Name          string    `json:"name"`
	BaseProductID entity.ID `json:"base_product_id"`
}

type CreateProduct struct {
	Name          string    `json:"name" binding:"required"`
	BaseProductID entity.ID `json:"base_product_id" binding:"required"`
}

type UpdateProduct struct {
	Name          string    `json:"name" binding:"required"`
	BaseProductID entity.ID `json:"base_product_id" binding:"required"`
}
