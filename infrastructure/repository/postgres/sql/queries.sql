-- Base products

-- name: GetBaseProductByID :one
SELECT * FROM base_products WHERE id = $1 LIMIT 1;

-- name: SearchBaseProduct :many
SELECT * FROM base_products WHERE name = $1;

-- name: ListBaseProduct :many
SELECT * FROM base_products ORDER BY name;

-- name: CreateBaseProduct :one
INSERT INTO base_products
(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateBaseProduct :exec
UPDATE base_products SET
(name, created_at, updated_at) = ($1, $2, $3)
WHERE id = $4;

-- name: DeleteBaseProduct :exec
DELETE FROM base_products WHERE id = $1;

-- ------------------------------------------------------------------------------------
-- Inventory

-- name: GetInventoryByID :one
SELECT * FROM inventories WHERE id = $1 LIMIT 1;

-- name: ListInventory :many
SELECT * FROM inventories ORDER BY created_at;

-- name: CreateInventory :one
INSERT INTO inventories
(id, user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateInventory :exec
UPDATE inventories SET
(user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at) = ($1, $2, $3, $4, $5, $6)
WHERE id = $7;

-- name: DeleteInventory :exec
DELETE FROM inventories WHERE id = $1;

-- ------------------------------------------------------------------------------------
-- Product

-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: SearchProduct :many
SELECT * FROM products WHERE name = $1;

-- name: ListProduct :many
SELECT * FROM products ORDER BY name;

-- name: CreateProduct :one
INSERT INTO products
(id, name, base_product_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateProduct :exec
UPDATE products SET
(name, base_product_id, created_at, updated_at) = ($1, $2, $3, $4)
WHERE id = $5;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- ------------------------------------------------------------------------------------
-- Unit of measure

-- name: GetUnitOfMeasureByID :one
SELECT * FROM units_of_measure WHERE id = $1 LIMIT 1;

-- name: SearchUnitOfMeasure :many
SELECT * FROM units_of_measure WHERE name = $1;

-- name: ListUnitOfMeasure :many
SELECT * FROM units_of_measure ORDER BY name;

-- name: CreateUnitOfMeasure :one
INSERT INTO units_of_measure
(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUnitOfMeasure :exec
UPDATE units_of_measure SET
(name, created_at, updated_at) = ($1, $2, $3)
WHERE id = $4;

-- name: DeleteUnitOfMeasure :exec
DELETE FROM units_of_measure WHERE id = $1;

-- ------------------------------------------------------------------------------------
-- Roles

-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = $1 LIMIT 1;

-- name: SearchRole :many
SELECT * FROM roles WHERE name = $1;

-- name: ListRole :many
SELECT * FROM roles ORDER BY name;

-- name: CreateRole :one
INSERT INTO roles
(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateRole :exec
UPDATE roles SET
(name, created_at, updated_at) = ($1, $2, $3)
WHERE id = $4;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;

-- ------------------------------------------------------------------------------------
-- Users

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: SearchUser :many
SELECT * FROM users WHERE first_name = $1;

-- name: ListUser :many
SELECT * FROM users ORDER BY first_name;

-- name: CreateUser :one
INSERT INTO users
(id, first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET
(first_name, last_name, address, phone, email, latitude, longitude, role_id, created_at, updated_at) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
WHERE id = $11;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- ------------------------------------------------------------------------------------
-- Customer

-- name: GetCustomerByID :one
SELECT * FROM customers WHERE id = $1 LIMIT 1;

-- name: SearchCustomer :many
SELECT * FROM customers WHERE first_name = $1;

-- name: ListCustomer :many
SELECT * FROM customers ORDER BY first_name;

-- name: CreateCustomer :one
INSERT INTO customers
(id, first_name, last_name, address, phone, email, latitude, longitude, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateCustomer :exec
UPDATE customers SET
(first_name, last_name, address, phone, email, latitude, longitude, created_at, updated_at) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
WHERE id = $10;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1;
