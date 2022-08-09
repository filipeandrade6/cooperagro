package postgres

import (
	"context"
	"errors"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"
	"github.com/jackc/pgx/v4"
)

func (r *Repo) GetInventoryByID(id entity.ID) (*entity.Inventory, error) {
	ctx := context.Background()
	i, err := r.db.GetInventoryByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entity.Inventory{
		ID:              i.ID,
		UserID:          i.UserID,
		ProductID:       i.ProductID,
		Quantity:        int(i.Quantity),
		UnitOfMeasureID: i.UnitOfMeasureID,
		CreatedAt:       i.CreatedAt,
		UpdatedAt:       i.UpdatedAt,
	}, nil
}

func (r *Repo) ListInventory() ([]*entity.Inventory, error) {
	ctx := context.Background()
	inventories, err := r.db.ListInventory(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var inventoriesOut []*entity.Inventory
	for _, inventory := range inventories {
		inventoriesOut = append(inventoriesOut, &entity.Inventory{
			ID:              inventory.ID,
			UserID:          inventory.UserID,
			ProductID:       inventory.ProductID,
			Quantity:        int(inventory.Quantity),
			UnitOfMeasureID: inventory.UnitOfMeasureID,
			CreatedAt:       inventory.CreatedAt,
			UpdatedAt:       inventory.UpdatedAt,
		})
	}

	return inventoriesOut, nil
}

func (r *Repo) CreateInventory(e *entity.Inventory) (entity.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateInventory(ctx, data.CreateInventoryParams{
		ID:              e.ID,
		UserID:          e.UserID, // TODO uns estao User outros User, decidir
		ProductID:       e.ProductID,
		Quantity:        int32(e.Quantity),
		UnitOfMeasureID: e.UnitOfMeasureID,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	})
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *Repo) UpdateInventory(e *entity.Inventory) error {
	ctx := context.Background()
	err := r.db.UpdateInventory(ctx, data.UpdateInventoryParams{
		UserID:          e.UserID,
		ProductID:       e.ProductID,
		Quantity:        int32(e.Quantity),
		UnitOfMeasureID: e.UnitOfMeasureID,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		ID:              e.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteInventory(id entity.ID) error {
	ctx := context.Background()
	err := r.db.DeleteInventory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
