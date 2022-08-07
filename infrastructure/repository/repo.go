package repository

import (
	"github.com/filipeandrade6/cooperagro/domain/usecases/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecases/customer"
	"github.com/filipeandrade6/cooperagro/domain/usecases/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecases/product"
	"github.com/filipeandrade6/cooperagro/domain/usecases/unitofmeasure"
)

type Repository interface {
	baseproduct.Repository
	customer.Repository
	inventory.Repository
	product.Repository
	unitofmeasure.Repository
}
