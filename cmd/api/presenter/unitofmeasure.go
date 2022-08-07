package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type UnitOfMeasure struct {
	ID   entities.ID `json:"id"`
	Name string      `json:"name"`
}
