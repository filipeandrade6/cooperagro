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

func NewPostgresRepo(urlConn string) (*Repo, error) {
	conn, err := pgxpool.Connect(context.Background(), urlConn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	// TODO dar o close com a aplicação

	r := New(conn)

	return &Repo{r}, nil
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
		return e.ID, err
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

// -- User

func (r *Repo) GetUserByID(id entities.ID) (*entities.User, error) {
	ctx := context.Background()
	c, err := r.db.GetUserByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		Phone:     c.Phone,
		Email:     c.Email,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
		Roles:     c.Roles,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (r *Repo) SearchUser(query string) ([]*entities.User, error) {
	ctx := context.Background()
	users, err := r.db.SearchUser(ctx, query) // TODO SearchUser faz busca na coluna first_name -> alterar depois
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var usersOut []*entities.User
	for _, user := range users {
		usersOut = append(usersOut, &entities.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Address:   user.Address,
			Phone:     user.Phone,
			Email:     user.Email,
			Latitude:  user.Latitude,
			Longitude: user.Longitude,
			Roles:     user.Roles,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return usersOut, nil
}

func (r *Repo) ListUser() ([]*entities.User, error) {
	ctx := context.Background()
	users, err := r.db.ListUser(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var usersOut []*entities.User
	for _, User := range users {
		usersOut = append(usersOut, &entities.User{
			ID:        User.ID,
			FirstName: User.FirstName,
			LastName:  User.LastName,
			Address:   User.Address,
			Phone:     User.Phone,
			Email:     User.Email,
			Latitude:  User.Latitude,
			Longitude: User.Longitude,
			Roles:     User.Roles,
			CreatedAt: User.CreatedAt,
			UpdatedAt: User.UpdatedAt,
		})
	}

	return usersOut, nil
}

func (r *Repo) CreateUser(e *entities.User) (entities.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateUser(ctx, CreateUserParams{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Address:   e.Address,
		Phone:     e.Phone,
		Email:     e.Email,
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Roles:     e.Roles,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	})
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *Repo) UpdateUser(e *entities.User) error {
	ctx := context.Background()
	err := r.db.UpdateUser(ctx, UpdateUserParams{
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Address:   e.Address,
		Phone:     e.Phone,
		Email:     e.Email,
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Roles:     e.Roles,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		ID:        e.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteUser(id entities.ID) error {
	ctx := context.Background()
	err := r.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// -- Inventory

func (r *Repo) GetInventoryByID(id entities.ID) (*entities.Inventory, error) {
	ctx := context.Background()
	i, err := r.db.GetInventoryByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.Inventory{
		ID:              i.ID,
		UserID:          i.UserID,
		ProductID:       i.ProductID,
		Quantity:        int(i.Quantity),
		UnitOfMeasureID: i.UnitOfMeasureID,
		CreatedAt:       i.CreatedAt,
		UpdatedAt:       i.UpdatedAt,
	}, nil
}

func (r *Repo) ListInventory() ([]*entities.Inventory, error) {
	ctx := context.Background()
	inventories, err := r.db.ListInventory(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var inventoriesOut []*entities.Inventory
	for _, inventory := range inventories {
		inventoriesOut = append(inventoriesOut, &entities.Inventory{
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

func (r *Repo) CreateInventory(e *entities.Inventory) (entities.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateInventory(ctx, CreateInventoryParams{
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

func (r *Repo) UpdateInventory(e *entities.Inventory) error {
	ctx := context.Background()
	err := r.db.UpdateInventory(ctx, UpdateInventoryParams{
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

func (r *Repo) DeleteInventory(id entities.ID) error {
	ctx := context.Background()
	err := r.db.DeleteInventory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// -- Product

func (r *Repo) GetProductByID(id entities.ID) (*entities.Product, error) {
	ctx := context.Background()
	p, err := r.db.GetProductByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.Product{
		ID:            p.ID,
		Name:          p.Name,
		BaseProductID: p.BaseProductID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}, nil
}

func (r *Repo) SearchProduct(query string) ([]*entities.Product, error) {
	ctx := context.Background()
	products, err := r.db.SearchProduct(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var productsOut []*entities.Product
	for _, product := range products {
		productsOut = append(productsOut, &entities.Product{
			ID:            product.ID,
			Name:          product.Name,
			BaseProductID: product.BaseProductID,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
		})
	}

	return productsOut, nil
}

func (r *Repo) ListProduct() ([]*entities.Product, error) {
	ctx := context.Background()
	products, err := r.db.ListProduct(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var productsOut []*entities.Product
	for _, product := range products {
		productsOut = append(productsOut, &entities.Product{
			ID:            product.ID,
			Name:          product.Name,
			BaseProductID: product.BaseProductID,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
		})
	}

	return productsOut, nil
}

func (r *Repo) CreateProduct(e *entities.Product) (entities.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateProduct(ctx, CreateProductParams{
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

func (r *Repo) UpdateProduct(e *entities.Product) error {
	ctx := context.Background()
	err := r.db.UpdateProduct(ctx, UpdateProductParams{
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

func (r *Repo) DeleteProduct(id entities.ID) error {
	ctx := context.Background()
	err := r.db.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// -- Unit of measure

func (r *Repo) GetUnitOfMeasureByID(id entities.ID) (*entities.UnitOfMeasure, error) {
	ctx := context.Background()
	u, err := r.db.GetUnitOfMeasureByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.UnitOfMeasure{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *Repo) SearchUnitOfMeasure(query string) ([]*entities.UnitOfMeasure, error) {
	ctx := context.Background()
	units, err := r.db.SearchUnitOfMeasure(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var unitsOut []*entities.UnitOfMeasure
	for _, unit := range units {
		unitsOut = append(unitsOut, &entities.UnitOfMeasure{
			ID:        unit.ID,
			Name:      unit.Name,
			CreatedAt: unit.CreatedAt,
			UpdatedAt: unit.UpdatedAt,
		})
	}

	return unitsOut, nil
}

func (r *Repo) ListUnitOfMeasure() ([]*entities.UnitOfMeasure, error) {
	ctx := context.Background()
	units, err := r.db.ListUnitOfMeasure(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var unitsOut []*entities.UnitOfMeasure
	for _, unit := range units {
		unitsOut = append(unitsOut, &entities.UnitOfMeasure{
			ID:        unit.ID,
			Name:      unit.Name,
			CreatedAt: unit.CreatedAt,
			UpdatedAt: unit.UpdatedAt,
		})
	}

	return unitsOut, nil
}

func (r *Repo) CreateUnitOfMeasure(e *entities.UnitOfMeasure) (entities.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateUnitOfMeasure(ctx, CreateUnitOfMeasureParams{
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

func (r *Repo) UpdateUnitOfMeasure(e *entities.UnitOfMeasure) error {
	ctx := context.Background()
	err := r.db.UpdateUnitOfMeasure(ctx, UpdateUnitOfMeasureParams{
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

func (r *Repo) DeleteUnitOfMeasure(id entities.ID) error {
	ctx := context.Background()
	err := r.db.DeleteUnitOfMeasure(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
