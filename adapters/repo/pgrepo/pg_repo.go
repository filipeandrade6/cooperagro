package pgrepo

import (
	"context"

	"github.com/filipeandrade6/cooperagro/adapters/log"
	"github.com/filipeandrade6/cooperagro/domain"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

// Repo implements the repo interface by using the ksql database.
type Repo struct {
	db ksql.Provider
}

// New instantiates a new Repo
func New(ctx context.Context, postgresURL string) (Repo, error) {
	db, err := kpgx.New(ctx, postgresURL, ksql.Config{})
	if err != nil {
		return Repo{}, domain.InternalErr("unable to start database", log.Body{
			"error": err.Error(),
		})
	}

	return Repo{
		db: db,
	}, nil
}

// TODO arrumar a documentação

// ============================================================================
// User

// ChangeUserEmail implements the repo.Users interface
func (r Repo) ChangeUserEmail(ctx context.Context, userID int, newEmail string) error {
	return changeUserEmail(ctx, r.db, userID, newEmail)
}

// UpsertUser implements the repo.Users interface
func (r Repo) UpsertUser(ctx context.Context, user domain.User) (userID int, _ error) {
	return upsertUser(ctx, r.db, user)
}

// GetUserByID implements the repo.Users interface
func (r Repo) GetUserByID(ctx context.Context, userID int) (domain.User, error) {
	return getUserByID(ctx, r.db, userID)
}

// GetUserByEmail implements the repo.Users interface
func (r Repo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return getUserByEmail(ctx, r.db, email)
}

// ============================================================================
// Product

// ChangeProductName implements the repo.Users interface
func (r Repo) ChangeProductName(ctx context.Context, userID int, newProductName string) error {
	return changeProductName(ctx, r.db, userID, newProductName)
}

// UpsertProduct implements the repo.Users interface
func (r Repo) UpsertProduct(ctx context.Context, product domain.Product) (userID int, _ error) {
	return upsertProduct(ctx, r.db, product)
}

// GetProductByID implements the repo.Users interface
func (r Repo) GetProductByID(ctx context.Context, productID int) (domain.Product, error) {
	return getProductByID(ctx, r.db, productID)
}

// GetProductByName implements the repo.Users interface
func (r Repo) GetProductByName(ctx context.Context, name string) (domain.Product, error) {
	return getProductByName(ctx, r.db, name)
}

// ============================================================================
// Base product

// ChangeBaseProductName implements the repo.Users interface
func (r Repo) ChangeBaseProductName(ctx context.Context, userID int, newBaseProductName string) error {
	return changeBaseProductName(ctx, r.db, userID, newBaseProductName)
}

// UpsertBaseProduct implements the repo.Users interface
func (r Repo) UpsertBaseProduct(ctx context.Context, baseProduct domain.BaseProduct) (baseProductID int, _ error) {
	return upsertBaseProduct(ctx, r.db, baseProduct)
}

// GetBaseProductByID implements the repo.Users interface
func (r Repo) GetBaseProductByID(ctx context.Context, userID int) (domain.BaseProduct, error) {
	return getBaseProductByID(ctx, r.db, userID)
}

// GetBaseProductByName implements the repo.Users interface
func (r Repo) GetBaseProductByName(ctx context.Context, name string) (domain.BaseProduct, error) {
	return getBaseProductByName(ctx, r.db, name)
}

// ============================================================================
// Role

// ChangeRoleName implements the repo.Users interface
func (r Repo) ChangeRoleName(ctx context.Context, userID int, newRoleName string) error {
	return changeRoleName(ctx, r.db, userID, newRoleName)
}

// UpsertRole implements the repo.Users interface
func (r Repo) UpsertRole(ctx context.Context, role domain.Role) (roleID int, _ error) {
	return upsertRole(ctx, r.db, role)
}

// GetRoleByID implements the repo.Users interface
func (r Repo) GetRoleByID(ctx context.Context, roleID int) (domain.Role, error) {
	return getRoleByID(ctx, r.db, roleID)
}

// GetRoleByName implements the repo.Users interface
func (r Repo) GetRoleByName(ctx context.Context, name string) (domain.Role, error) {
	return getRoleByName(ctx, r.db, name)
}

// ============================================================================
// Unit of measure

// ChangeUnitOfMeasureName implements the repo.Users interface
func (r Repo) ChangeUnitOfMeasureName(ctx context.Context, userID int, newUnitOfMeasureName string) error {
	return changeUnitOfMeasureName(ctx, r.db, userID, newUnitOfMeasureName)
}

// UpsertUnitOfMeasure implements the repo.Users interface
func (r Repo) UpsertUnitOfMeasure(ctx context.Context, unitOfMeasure domain.UnitOfMeasure) (unitOfMeasureID int, _ error) {
	return upsertUnitOfMeasure(ctx, r.db, unitOfMeasure)
}

// GetUnitOfMeasureByID implements the repo.Users interface
func (r Repo) GetUnitOfMeasureByID(ctx context.Context, unitOfMeasureID int) (domain.UnitOfMeasure, error) {
	return getUnitOfMeasureByID(ctx, r.db, unitOfMeasureID)
}

// GetUnitOfMeasureByName implements the repo.Users interface
func (r Repo) GetUnitOfMeasureByName(ctx context.Context, name string) (domain.UnitOfMeasure, error) {
	return getUnitOfMeasureByName(ctx, r.db, name)
}
