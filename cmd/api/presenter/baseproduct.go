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

type EchoBaseProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EchoCreateBaseProduct struct {
	Name string `form:"name" json:"name"`
}

type EchoUpdateBaseProduct struct {
	ID   string `json:"id"`
	Name string `form:"name" json:"name"`
}

type EchoDeleteBaseProduct struct {
	ID string `json:"id"`
}

// Juntar todos?

// TODO escolher qual for utilizar e deletar o outro e remover prefixo
