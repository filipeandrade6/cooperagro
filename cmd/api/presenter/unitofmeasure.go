package presenter

import "github.com/filipeandrade6/cooperagro/domain/entity"

type UnitOfMeasure struct {
	ID   entity.ID `json:"id"`
	Name string    `json:"name"`
}

type CreateUnitOfMeasure struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUnitOfMeasure struct {
	Name string `json:"name" binding:"required"`
}
