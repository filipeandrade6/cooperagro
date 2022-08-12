package repository

import (
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecase/product"
	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
)

type Repository interface {
	baseproduct.Repository
	inventory.Repository
	product.Repository
	unitofmeasure.Repository
	user.Repository
}
