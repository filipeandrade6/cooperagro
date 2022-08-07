package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/filipeandrade6/cooperagro/domain/entities"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO adicionar logger

type Repo struct {
	db *Queries
}

func NewPostgresRepo(urlConn string) (*Queries, error) {
	conn, err := pgxpool.Connect(context.Background(), urlConn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	// TODO dar o close com a aplicação

	r := New(conn)

	return r, nil
}

// -- Base product

func (r *Repo) GetBaseProductByID(id entities.ID) (*entities.BaseProduct, error) {
	ctx := context.Background()
	bp, err := r.db.GetBaseProductByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.BaseProduct{
		ID:        bp.ID,
		Name:      bp.Name,
		CreatedAt: bp.CreatedAt,
		UpdatedAt: bp.UpdatedAt,
	}, nil
}

func (r *Repo) SearchBaseProduct(query string) ([]*entities.BaseProduct, error) {
	ctx := context.Background()
	bps, err := r.db.SearchBaseProduct(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var bpsOut []*entities.BaseProduct

	for _, bp := range bps {
		bpsOut = append(bpsOut, &entities.BaseProduct{
			ID:        bp.ID,
			Name:      bp.Name,
			CreatedAt: bp.CreatedAt,
			UpdatedAt: bp.UpdatedAt,
		})
	}

	return bpsOut, nil
}

func (r *Repo) ListBaseProduct() ([]*entities.BaseProduct, error) {
	ctx := context.Background()
	bps, err := r.db.ListBaseProduct(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var bpsOut []*entities.BaseProduct

	for _, bp := range bps {
		bpsOut = append(bpsOut, &entities.BaseProduct{
			ID:        bp.ID,
			Name:      bp.Name,
			CreatedAt: bp.CreatedAt,
			UpdatedAt: bp.UpdatedAt,
		})
	}

	return bpsOut, nil
}

func (r *Repo) CreateBaseProduct(e *entities.BaseProduct) (entities.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateBaseProduct(ctx, CreateBaseProductParams{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	})
	if err != nil {
		return entities.NewID(), err
	}

	return e.ID, nil
}

func (r *Repo) UpdateBaseProduct(e *entities.BaseProduct) error {
	ctx := context.Background()
	err := r.db.UpdateBaseProduct(ctx, UpdateBaseProductParams{
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

func (r *Repo) DeleteBaseProduct(id entities.ID) error {
	ctx := context.Background()
	err := r.db.DeleteBaseProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// -- Customer

func (r *Repo) GetCustomerByID(id entities.ID) (*entities.Customer, error) {
	ctx := context.Background()
	c, err := r.db.GetCustomerByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.Customer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		Phone:     c.Phone,
		Email:     c.Email,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (r *Repo) SearchCustomer(query string) ([]*entities.Customer, error) {
	return nil, nil
}

func (r *Repo) ListCustomer() ([]*entities.Customer, error) {
	return nil, nil
}

func (r *Repo) CreateCustomer(e *entities.Customer) (entities.ID, error) {
	return entities.NewID(), nil
}

func (r *Repo) UpdateCustomer(e *entities.Customer) error {
	return nil
}

func (r *Repo) DeleteCustomer(id entities.ID) error {
	return nil
}

// -- Inventory

func (r *Repo) GetInventoryByID(id entities.ID) (*entities.Inventory, error) {
	return nil, nil
}

func (r *Repo) ListInventory() ([]*entities.Inventory, error) {
	return nil, nil
}

func (r *Repo) CreateInventory(e *entities.Inventory) entities.ID {
	return entities.NewID()
}

func (r *Repo) UpdateInventory(e *entities.Inventory) error {
	return nil
}

func (r *Repo) DeleteInventory(id entities.ID) error {
	return nil
}

// -- Product

func (r *Repo) GetProductByID(id entities.ID) (*entities.Product, error) {
	return nil, nil
}

func (r *Repo) SearchProduct(query string) ([]*entities.Product, error) {
	return nil, nil
}

func (r *Repo) ListProduct() ([]*entities.Product, error) {
	return nil, nil
}

func (r *Repo) CreateProduct(e *entities.Product) (entities.ID, error) {
	return entities.NewID(), nil
}

func (r *Repo) UpdateProduct(e *entities.Product) error {
	return nil
}

func (r *Repo) DeleteProduct(id entities.ID) error {
	return nil
}

// -- Unit of measure

func (r *Repo) GetUnitOfMeasureByID(id entities.ID) (*entities.UnitOfMeasure, error) {
	return nil, nil
}

func (r *Repo) SearchUnitOfMeasure(query string) ([]*entities.UnitOfMeasure, error) {
	return nil, nil
}

func (r *Repo) ListUnitOfMeasure() ([]*entities.UnitOfMeasure, error) {
	return nil, nil
}

func (r *Repo) CreateUnitOfMeasure(e *entities.UnitOfMeasure) (entities.ID, error) {
	return entities.NewID(), nil
}

func (r *Repo) UpdateUnitOfMeasure(e *entities.UnitOfMeasure) error {
	return nil
}

func (r *Repo) DeleteUnitOfMeasure(id entities.ID) error {
	return nil
}
