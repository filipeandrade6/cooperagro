package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type BaseProduct struct {
	ID   entities.ID `json:"id"`
	Name string      `json:"name"`
}

type CreateBaseProduct struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdateBaseProduct struct {
	Name string `form:"name" json:"name" binding:"required"`
}
