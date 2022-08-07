package presenter

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Customer struct {
	ID        entities.ID `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Address   string      `json:"address"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Latitude  float32     `json:"latitude"`
	Longitude float32     `json:"longitude"`
}

type CreateCustomer struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Phone     string  `json:"phone" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type UpdateCustomer struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Phone     string  `json:"phone" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}
