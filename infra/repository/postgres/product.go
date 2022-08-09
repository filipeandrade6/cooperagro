package postgres

import (
	"context"
	"errors"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"
	"github.com/jackc/pgx/v4"
)

func (r *Repo) GetProductByID(id entity.ID) (*entity.Product, error) {
	ctx := context.Background()
	p, err := r.db.GetProductByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entity.Product{
		ID:            p.ID,
		Name:          p.Name,
		BaseProductID: p.BaseProductID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}, nil
}

func (r *Repo) SearchProduct(query string) ([]*entity.Product, error) {
	ctx := context.Background()
	products, err := r.db.SearchProduct(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var productsOut []*entity.Product
	for _, product := range products {
		productsOut = append(productsOut, &entity.Product{
			ID:            product.ID,
			Name:          product.Name,
			BaseProductID: product.BaseProductID,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
		})
	}

	return productsOut, nil
}

func (r *Repo) ListProduct() ([]*entity.Product, error) {
	ctx := context.Background()
	products, err := r.db.ListProduct(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var productsOut []*entity.Product
	for _, product := range products {
		productsOut = append(productsOut, &entity.Product{
			ID:            product.ID,
			Name:          product.Name,
			BaseProductID: product.BaseProductID,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
		})
	}

	return productsOut, nil
}

func (r *Repo) CreateProduct(e *entity.Product) (entity.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateProduct(ctx, data.CreateProductParams{
		ID:            e.ID,
		Name:          e.Name,
		BaseProductID: e.BaseProductID,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	})
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *Repo) UpdateProduct(e *entity.Product) error {
	ctx := context.Background()
	err := r.db.UpdateProduct(ctx, data.UpdateProductParams{
		Name:          e.Name,
		BaseProductID: e.BaseProductID,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		ID:            e.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteProduct(id entity.ID) error {
	ctx := context.Background()
	err := r.db.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
