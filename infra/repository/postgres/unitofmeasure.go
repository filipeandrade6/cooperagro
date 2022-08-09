package postgres

import (
	"context"
	"errors"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"
	"github.com/jackc/pgx/v4"
)

func (r *Repo) GetUnitOfMeasureByID(id entity.ID) (*entity.UnitOfMeasure, error) {
	ctx := context.Background()
	u, err := r.db.GetUnitOfMeasureByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entity.UnitOfMeasure{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *Repo) SearchUnitOfMeasure(query string) ([]*entity.UnitOfMeasure, error) {
	ctx := context.Background()
	units, err := r.db.SearchUnitOfMeasure(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var unitsOut []*entity.UnitOfMeasure
	for _, unit := range units {
		unitsOut = append(unitsOut, &entity.UnitOfMeasure{
			ID:        unit.ID,
			Name:      unit.Name,
			CreatedAt: unit.CreatedAt,
			UpdatedAt: unit.UpdatedAt,
		})
	}

	return unitsOut, nil
}

func (r *Repo) ListUnitOfMeasure() ([]*entity.UnitOfMeasure, error) {
	ctx := context.Background()
	units, err := r.db.ListUnitOfMeasure(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var unitsOut []*entity.UnitOfMeasure
	for _, unit := range units {
		unitsOut = append(unitsOut, &entity.UnitOfMeasure{
			ID:        unit.ID,
			Name:      unit.Name,
			CreatedAt: unit.CreatedAt,
			UpdatedAt: unit.UpdatedAt,
		})
	}

	return unitsOut, nil
}

func (r *Repo) CreateUnitOfMeasure(e *entity.UnitOfMeasure) (entity.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateUnitOfMeasure(ctx, data.CreateUnitOfMeasureParams{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	})
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *Repo) UpdateUnitOfMeasure(e *entity.UnitOfMeasure) error {
	ctx := context.Background()
	err := r.db.UpdateUnitOfMeasure(ctx, data.UpdateUnitOfMeasureParams{
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		ID:        e.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteUnitOfMeasure(id entity.ID) error {
	ctx := context.Background()
	err := r.db.DeleteUnitOfMeasure(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
