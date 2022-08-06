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
(name, created_at, updated_at) = ($1, $2, $3, $4)
WHERE id = $5;

-- name: DeleteBaseProduct :exec
DELETE FROM base_products WHERE id = $1;






-- name: GetBaseProductByID :one
-- name: SearchBaseProduct :many
-- name: ListBaseProduct :many
-- name: CreateBaseProduct :one
-- name: UpdateBaseProduct :exec
-- name: DeleteBaseProduct :exec