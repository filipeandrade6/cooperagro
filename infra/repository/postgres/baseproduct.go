package postgres

import (
	"context"
	"errors"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// -- Base product

func (r *Repo) GetBaseProductByID(id entity.ID) (*entity.BaseProduct, error) {
	ctx := context.Background()
	bp, err := r.db.GetBaseProductByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entity.BaseProduct{
		ID:        bp.ID,
		Name:      bp.Name,
		CreatedAt: bp.CreatedAt,
		UpdatedAt: bp.UpdatedAt,
	}, nil
}

func (r *Repo) SearchBaseProduct(query string) ([]*entity.BaseProduct, error) {
	ctx := context.Background()
	bps, err := r.db.SearchBaseProduct(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var bpsOut []*entity.BaseProduct

	for _, bp := range bps {
		bpsOut = append(bpsOut, &entity.BaseProduct{
			ID:        bp.ID,
			Name:      bp.Name,
			CreatedAt: bp.CreatedAt,
			UpdatedAt: bp.UpdatedAt,
		})
	}

	return bpsOut, nil
}

func (r *Repo) ListBaseProduct() ([]*entity.BaseProduct, error) {
	ctx := context.Background()
	bps, err := r.db.ListBaseProduct(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var bpsOut []*entity.BaseProduct

	for _, bp := range bps {
		bpsOut = append(bpsOut, &entity.BaseProduct{
			ID:        bp.ID,
			Name:      bp.Name,
			CreatedAt: bp.CreatedAt,
			UpdatedAt: bp.UpdatedAt,
		})
	}

	return bpsOut, nil
}

func (r *Repo) CreateBaseProduct(e *entity.BaseProduct) (entity.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateBaseProduct(ctx, data.CreateBaseProductParams{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	})

	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // TODO verificar esse tratamento de erro
		return e.ID, entity.ErrEntityAlreadyExists
	}

	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *Repo) UpdateBaseProduct(e *entity.BaseProduct) error {
	ctx := context.Background()
	err := r.db.UpdateBaseProduct(ctx, data.UpdateBaseProductParams{
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		ID:        e.ID,
	})

	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // TODO verificar esse tratamento de erro
		return entity.ErrEntityAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteBaseProduct(id entity.ID) error {
	ctx := context.Background()
	err := r.db.DeleteBaseProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
