package presenter

import "github.com/filipeandrade6/cooperagro/domain/entity"

type BaseProduct struct {
	ID   entity.ID `json:"id"`
	Name string    `json:"name"`
}

type CreateBaseProduct struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdateBaseProduct struct {
	Name string `form:"name" json:"name" binding:"required"`
}
