package repo

import (
	"context"

	"github.com/filipeandrade6/cooperagro/domain"
)

type Repo interface {
	UpsertUser(ctx context.Context, user domain.User) (userID int, err error)
	GetUserByID(ctx context.Context, userID int) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)

	UpsertProduct(ctx context.Context, product domain.Product) (productID int, err error)
	GetProductByID(ctx context.Context, productID int) (domain.Product, error)
	GetProductByName(ctx context.Context, name string) (domain.Product, error)

	UpsertBaseProduct(ctx context.Context, baseProduct domain.BaseProduct) (baseProductID int, err error)
	GetBaseProductByID(ctx context.Context, baseProductID int) (domain.BaseProduct, error)
	GetBaseProductByName(ctx context.Context, name string) (domain.BaseProduct, error)

	UpsertRole(ctx context.Context, role domain.Role) (roleID int, err error)
	GetRoleByID(ctx context.Context, roleID int) (domain.Role, error)
	GetRoleByName(ctx context.Context, name string) (domain.Role, error)

	UpsertUnitOfMeasure(ctx context.Context, unitOfMeasure domain.UnitOfMeasure) (unitOfMeasureID int, err error)
	GetUnitOfMeasureByID(ctx context.Context, unitOfMeasureID int) (domain.UnitOfMeasure, error)
	GetUnitOfMeasureByName(ctx context.Context, name string) (domain.UnitOfMeasure, error)
}

// Users ... TODO adicionar
type Users interface {
	UpsertUser(ctx context.Context, user domain.User) (userID int, err error)
	GetUserByID(ctx context.Context, userID int) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

// Products ... TODO adicionar
type Products interface {
	UpsertProduct(ctx context.Context, product domain.Product) (productID int, err error)
	GetProductByID(ctx context.Context, productID int) (domain.Product, error)
	GetProductByName(ctx context.Context, name string) (domain.Product, error)
}

// BaseProducts ... TODO adicionar
type BaseProducts interface {
	UpsertBaseProduct(ctx context.Context, baseProduct domain.BaseProduct) (baseProductID int, err error)
	GetBaseProductByID(ctx context.Context, baseProductID int) (domain.BaseProduct, error)
	GetBaseProductByName(ctx context.Context, name string) (domain.BaseProduct, error)
}

// Roles ... TODO adicionar
type Roles interface {
	UpsertRole(ctx context.Context, role domain.Role) (roleID int, err error)
	GetRoleByID(ctx context.Context, roleID int) (domain.Role, error)
	GetRoleByName(ctx context.Context, name string) (domain.Role, error)
}

// UnitsOfMeasure ... TODO adicionar
type UnitsOfMeasure interface {
	UpsertUnitOfMeasure(ctx context.Context, unitOfMeasure domain.UnitOfMeasure) (unitOfMeasureID int, err error)
	GetUnitOfMeasureByID(ctx context.Context, unitOfMeasureID int) (domain.UnitOfMeasure, error)
	GetUnitOfMeasureByName(ctx context.Context, name string) (domain.UnitOfMeasure, error)
}
