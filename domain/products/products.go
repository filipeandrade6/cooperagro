package products

import (
	"github.com/filipeandrade6/cooperagro-go/adapters/log"
	"github.com/filipeandrade6/cooperagro-go/adapters/repo"
)

type Service struct {
	logger       log.Provider
	productsRepo repo.Products
}
