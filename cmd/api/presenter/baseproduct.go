package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type BaseProduct struct {
	ID   entities.ID `json:"id"`
	Name string      `json:"name"`
}
