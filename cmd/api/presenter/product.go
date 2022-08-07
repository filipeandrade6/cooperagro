package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Product struct {
	ID          entities.ID `json:"id"`
	Name        string      `json:"name"`
	BaseProduct entities.ID `json:"base_product"`
}
