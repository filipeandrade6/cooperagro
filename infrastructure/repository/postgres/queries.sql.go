// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: queries.sql

package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createBaseProduct = `-- name: CreateBaseProduct :one
INSERT INTO base_products
(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING id, name, created_at, updated_at
`

type CreateBaseProductParams struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateBaseProduct(ctx context.Context, arg CreateBaseProductParams) (BaseProduct, error) {
	row := q.db.QueryRow(ctx, createBaseProduct,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i BaseProduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createInventory = `-- name: CreateInventory :one
INSERT INTO inventories
(id, user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at
`

type CreateInventoryParams struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	ProductID       uuid.UUID
	Quantity        int32
	UnitOfMeasureID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (q *Queries) CreateInventory(ctx context.Context, arg CreateInventoryParams) (Inventory, error) {
	row := q.db.QueryRow(ctx, createInventory,
		arg.ID,
		arg.UserID,
		arg.ProductID,
		arg.Quantity,
		arg.UnitOfMeasureID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.Quantity,
		&i.UnitOfMeasureID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
INSERT INTO products
(id, name, base_product_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, base_product_id, created_at, updated_at
`

type CreateProductParams struct {
	ID            interface{}
	Name          string
	BaseProductID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.BaseProductID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BaseProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles
(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING id, name, created_at, updated_at
`

type CreateRoleParams struct {
	ID        uuid.UUID
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUnitOfMeasure = `-- name: CreateUnitOfMeasure :one
INSERT INTO units_of_measure
(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING id, name, created_at, updated_at
`

type CreateUnitOfMeasureParams struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUnitOfMeasure(ctx context.Context, arg CreateUnitOfMeasureParams) (UnitsOfMeasure, error) {
	row := q.db.QueryRow(ctx, createUnitOfMeasure,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i UnitsOfMeasure
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users
(id, first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Email     string
	Latitude  pgtype.Numeric
	Longitude pgtype.Numeric
	RoleID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Address,
		arg.Phone,
		arg.Email,
		arg.Latitude,
		arg.Longitude,
		arg.RoleID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Address,
		&i.Phone,
		&i.Email,
		&i.Latitude,
		&i.Longitude,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBaseProduct = `-- name: DeleteBaseProduct :exec
DELETE FROM base_products WHERE id = $1
`

func (q *Queries) DeleteBaseProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteBaseProduct, id)
	return err
}

const deleteInventory = `-- name: DeleteInventory :exec
DELETE FROM inventories WHERE id = $1
`

func (q *Queries) DeleteInventory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteInventory, id)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id interface{}) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteRole, id)
	return err
}

const deleteUnitOfMeasure = `-- name: DeleteUnitOfMeasure :exec
DELETE FROM units_of_measure WHERE id = $1
`

func (q *Queries) DeleteUnitOfMeasure(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUnitOfMeasure, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getBaseProductByID = `-- name: GetBaseProductByID :one

SELECT id, name, created_at, updated_at FROM base_products WHERE id = $1 LIMIT 1
`

// Base products
func (q *Queries) GetBaseProductByID(ctx context.Context, id uuid.UUID) (BaseProduct, error) {
	row := q.db.QueryRow(ctx, getBaseProductByID, id)
	var i BaseProduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInventoryByID = `-- name: GetInventoryByID :one

SELECT id, user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at FROM inventories WHERE id = $1 LIMIT 1
`

// ------------------------------------------------------------------------------------
// Inventory
func (q *Queries) GetInventoryByID(ctx context.Context, id uuid.UUID) (Inventory, error) {
	row := q.db.QueryRow(ctx, getInventoryByID, id)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.Quantity,
		&i.UnitOfMeasureID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductByID = `-- name: GetProductByID :one

SELECT id, name, base_product_id, created_at, updated_at FROM products WHERE id = $1 LIMIT 1
`

// ------------------------------------------------------------------------------------
// Product
func (q *Queries) GetProductByID(ctx context.Context, id interface{}) (Product, error) {
	row := q.db.QueryRow(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BaseProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoleByID = `-- name: GetRoleByID :one

SELECT id, name, created_at, updated_at FROM roles WHERE id = $1 LIMIT 1
`

// ------------------------------------------------------------------------------------
// Roles
func (q *Queries) GetRoleByID(ctx context.Context, id uuid.UUID) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByID, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUnitOfMeasureByID = `-- name: GetUnitOfMeasureByID :one

SELECT id, name, created_at, updated_at FROM units_of_measure WHERE id = $1 LIMIT 1
`

// ------------------------------------------------------------------------------------
// Unit of measure
func (q *Queries) GetUnitOfMeasureByID(ctx context.Context, id uuid.UUID) (UnitsOfMeasure, error) {
	row := q.db.QueryRow(ctx, getUnitOfMeasureByID, id)
	var i UnitsOfMeasure
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one

SELECT id, first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

// ------------------------------------------------------------------------------------
// Users
func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Address,
		&i.Phone,
		&i.Email,
		&i.Latitude,
		&i.Longitude,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBaseProduct = `-- name: ListBaseProduct :many
SELECT id, name, created_at, updated_at FROM base_products ORDER BY name
`

func (q *Queries) ListBaseProduct(ctx context.Context) ([]BaseProduct, error) {
	rows, err := q.db.Query(ctx, listBaseProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BaseProduct
	for rows.Next() {
		var i BaseProduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listInventory = `-- name: ListInventory :many
SELECT id, user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at FROM inventories ORDER BY created_at
`

func (q *Queries) ListInventory(ctx context.Context) ([]Inventory, error) {
	rows, err := q.db.Query(ctx, listInventory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Inventory
	for rows.Next() {
		var i Inventory
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.Quantity,
			&i.UnitOfMeasureID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProduct = `-- name: ListProduct :many
SELECT id, name, base_product_id, created_at, updated_at FROM products ORDER BY name
`

func (q *Queries) ListProduct(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BaseProductID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRole = `-- name: ListRole :many
SELECT id, name, created_at, updated_at FROM roles ORDER BY name
`

func (q *Queries) ListRole(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, listRole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUnitOfMeasure = `-- name: ListUnitOfMeasure :many
SELECT id, name, created_at, updated_at FROM units_of_measure ORDER BY name
`

func (q *Queries) ListUnitOfMeasure(ctx context.Context) ([]UnitsOfMeasure, error) {
	rows, err := q.db.Query(ctx, listUnitOfMeasure)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnitsOfMeasure
	for rows.Next() {
		var i UnitsOfMeasure
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUser = `-- name: ListUser :many
SELECT id, first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at FROM users ORDER BY first_name
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Address,
			&i.Phone,
			&i.Email,
			&i.Latitude,
			&i.Longitude,
			&i.RoleID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchBaseProduct = `-- name: SearchBaseProduct :many
SELECT id, name, created_at, updated_at FROM base_products WHERE name = $1
`

func (q *Queries) SearchBaseProduct(ctx context.Context, name string) ([]BaseProduct, error) {
	rows, err := q.db.Query(ctx, searchBaseProduct, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BaseProduct
	for rows.Next() {
		var i BaseProduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchProduct = `-- name: SearchProduct :many
SELECT id, name, base_product_id, created_at, updated_at FROM products WHERE name = $1
`

func (q *Queries) SearchProduct(ctx context.Context, name string) ([]Product, error) {
	rows, err := q.db.Query(ctx, searchProduct, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BaseProductID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchRole = `-- name: SearchRole :many
SELECT id, name, created_at, updated_at FROM roles WHERE name = $1
`

func (q *Queries) SearchRole(ctx context.Context, name sql.NullString) ([]Role, error) {
	rows, err := q.db.Query(ctx, searchRole, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchUnitOfMeasure = `-- name: SearchUnitOfMeasure :many
SELECT id, name, created_at, updated_at FROM units_of_measure WHERE name = $1
`

func (q *Queries) SearchUnitOfMeasure(ctx context.Context, name string) ([]UnitsOfMeasure, error) {
	rows, err := q.db.Query(ctx, searchUnitOfMeasure, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnitsOfMeasure
	for rows.Next() {
		var i UnitsOfMeasure
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchUser = `-- name: SearchUser :many
SELECT id, first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at FROM users WHERE first_name = $1
`

func (q *Queries) SearchUser(ctx context.Context, firstName string) ([]User, error) {
	rows, err := q.db.Query(ctx, searchUser, firstName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Address,
			&i.Phone,
			&i.Email,
			&i.Latitude,
			&i.Longitude,
			&i.RoleID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBaseProduct = `-- name: UpdateBaseProduct :exec
UPDATE base_products SET
(name, created_at, updated_at) = ($1, $2, $3)
WHERE id = $4
`

type UpdateBaseProductParams struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateBaseProduct(ctx context.Context, arg UpdateBaseProductParams) error {
	_, err := q.db.Exec(ctx, updateBaseProduct,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateInventory = `-- name: UpdateInventory :exec
UPDATE inventories SET
(user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at) = ($1, $2, $3, $4, $5, $6)
WHERE id = $7
`

type UpdateInventoryParams struct {
	UserID          uuid.UUID
	ProductID       uuid.UUID
	Quantity        int32
	UnitOfMeasureID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ID              uuid.UUID
}

func (q *Queries) UpdateInventory(ctx context.Context, arg UpdateInventoryParams) error {
	_, err := q.db.Exec(ctx, updateInventory,
		arg.UserID,
		arg.ProductID,
		arg.Quantity,
		arg.UnitOfMeasureID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products SET
(name, base_product_id, created_at, updated_at) = ($1, $2, $3, $4)
WHERE id = $5
`

type UpdateProductParams struct {
	Name          string
	BaseProductID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ID            interface{}
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.Name,
		arg.BaseProductID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateRole = `-- name: UpdateRole :exec
UPDATE roles SET
(name, created_at, updated_at) = ($1, $2, $3)
WHERE id = $4
`

type UpdateRoleParams struct {
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.Exec(ctx, updateRole,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateUnitOfMeasure = `-- name: UpdateUnitOfMeasure :exec
UPDATE units_of_measure SET
(name, created_at, updated_at) = ($1, $2, $3)
WHERE id = $4
`

type UpdateUnitOfMeasureParams struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateUnitOfMeasure(ctx context.Context, arg UpdateUnitOfMeasureParams) error {
	_, err := q.db.Exec(ctx, updateUnitOfMeasure,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET
(first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
WHERE id = $11
`

type UpdateUserParams struct {
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Email     string
	Latitude  pgtype.Numeric
	Longitude pgtype.Numeric
	RoleID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.Address,
		arg.Phone,
		arg.Email,
		arg.Latitude,
		arg.Longitude,
		arg.RoleID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
