package repo

import (
	"context"

	"github.com/filipeandrade6/cooperagro/domain"
)

// Users represents the operations we use for
// retrieving a user from a persistent storage
type Users interface {
	GetUser(ctx context.Context, userID int) (domain.User, error)
	UpsertUser(ctx context.Context, user domain.User) (userID int, err error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

// ============================================================================
// Meu

type Products interface {
	GetProduct(ctx context.Context, productID int) (domain.Product, error)
	UpsertProduct(ctx context.Context, product domain.Product) (productID int, err error)
	GetProductByName(ctx context.Context, name string) (domain.Product, error)
}
