package pgrepo

import (
	"context"

	"github.com/filipeandrade6/cooperagro-go/adapters/log"
	"github.com/filipeandrade6/cooperagro-go/domain"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

// UsersRepo implements the repo.Users interface by using the ksql database.
type UsersRepo struct {
	db ksql.Provider
}

// New instantiates a new UsersRepo
func New(ctx context.Context, postgresURL string) (UsersRepo, error) {
	db, err := kpgx.New(ctx, postgresURL, ksql.Config{})
	if err != nil {
		return UsersRepo{}, domain.InternalErr("unable to start database", log.Body{
			"error": err.Error(),
		})
	}

	return UsersRepo{
		db: db,
	}, nil
}

// ChangeUserEmail implements the repo.Users interface
func (u UsersRepo) ChangeUserEmail(ctx context.Context, userID int, newEmail string) error {
	return changeUserEmail(ctx, u.db, userID, newEmail)
}

// UpsertUser implements the repo.Users interface
func (u UsersRepo) UpsertUser(ctx context.Context, user domain.User) (userID int, _ error) {
	return upsertUser(ctx, u.db, user)
}

// GetUser implements the repo.Users interface
func (u UsersRepo) GetUser(ctx context.Context, userID int) (domain.User, error) {
	return getUser(ctx, u.db, userID)
}

// GetUserByEmail implements the repo.Users interface
func (u UsersRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return getUserByEmail(ctx, u.db, email)
}

// ============================================================================
// Meu

type ProductsRepo struct {
	db ksql.Provider
}

// NewProducts instantiates a new UsersRepo
func NewProducts(ctx context.Context, postgresURL string) (ProductsRepo, error) {
	db, err := kpgx.New(ctx, postgresURL, ksql.Config{})
	if err != nil {
		return ProductsRepo{}, domain.InternalErr("unable to start database", log.Body{
			"error": err.Error(),
		})
	}

	return ProductsRepo{
		db: db,
	}, nil
}

// ChangeProductName implements the repo.Products interface
func (p ProductsRepo) ChangeProductName(ctx context.Context, productID int, newProductName string) error {
	return changeProductName(ctx, p.db, productID, newProductName)
}

// UpsertProduct implements the repo.Products interface
func (p ProductsRepo) UpsertProduct(ctx context.Context, product domain.Product) (productID int, _ error) {
	return upsertProduct(ctx, p.db, product)
}

// GetProduct implements the repo.Products interface
func (p ProductsRepo) GetProduct(ctx context.Context, productID int) (domain.Product, error) {
	return getProduct(ctx, p.db, productID)
}

// GetProductByName implements the repo.Products interface
func (p ProductsRepo) GetProductByName(ctx context.Context, email string) (domain.Product, error) {
	return getProductByName(ctx, p.db, email)
}
