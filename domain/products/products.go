package products

import (
	"context"

	"github.com/filipeandrade6/cooperagro-go/adapters/log"
	"github.com/filipeandrade6/cooperagro-go/adapters/repo"
	"github.com/filipeandrade6/cooperagro-go/domain"
)

// Usually its here where the business logic complexity builds up,
// but since this is just an example both these functions are actually
// very simple, but in real world scenarios you would want to make
// these contain all your business logic.

type Service struct {
	logger       log.Provider
	productsRepo repo.Products
}

func NewService(
	logger log.Provider,
	productsRepo repo.Products,
) Service {
	return Service{
		logger:       logger,
		productsRepo: productsRepo,
	}
}

func (s Service) UpsertProduct(ctx context.Context, product domain.Product) (productID int, _ error) {
	productID, err := s.productsRepo.UpsertProduct(ctx, product)
	if err != nil {
		return 0, err
	}

	s.logger.Info(ctx, "product created", log.Body{
		"product_id": productID,
	})

	return productID, nil
}

func (s Service) GetProduct(ctx context.Context, productID int) (domain.Product, error) {
	product, err := s.productsRepo.GetProduct(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
