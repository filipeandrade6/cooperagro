package repository

import (
	"github.com/filipeandrade6/cooperagro/domain/usecases/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecases/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecases/product"
	"github.com/filipeandrade6/cooperagro/domain/usecases/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/domain/usecases/user"
)

type Repository interface {
	baseproduct.Repository
	user.Repository
	inventory.Repository
	product.Repository
	unitofmeasure.Repository
}
