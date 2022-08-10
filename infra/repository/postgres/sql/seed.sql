-- User
INSERT INTO
users(id, first_name, last_name, address, phone, email, latitude, longitude, roles, password, created_at, updated_at)
VALUES ('6bb774bd-fc80-40a9-a063-c7838209ec54', 'Filipe', 'Andrade', 'Main street', '5561555554444', 'filipe@email.com', -12.123456, -12.123456, '{admin, producer}', '$2a$10$nTxInXic3WCz14l64ycdx.78LJxBNGcw4/yT4LkdD9WZmFwAy/.pW', NOW(), NOW());

INSERT INTO
users(id, first_name, last_name, address, phone, email, latitude, longitude, roles, password, created_at, updated_at)
VALUES ('ae9e3267-6df2-4b2d-af3b-a7bb21fbfd26', 'Ana', 'Andrade', 'Main street', '5561555554444', 'ana@email.com', -12.123456, -12.123456, '{producer}', '$2a$10$nTxInXic3WCz14l64ycdx.78LJxBNGcw4/yT4LkdD9WZmFwAy/.pW', NOW(), NOW());

-- Unit of measure
INSERT INTO
units_of_measure (id, name, created_at, updated_at)
VALUES ('d01637a9-d6a4-44c9-92dc-12f25363cbe9', 'kilogram', NOW(), NOW());

-- Base product
INSERT INTO
base_products (id, name, created_at, updated_at)
VALUES ('28fb96f1-f213-4b93-9379-493b298c879b', 'laranja', NOW(), NOW());

-- Product
INSERT INTO
products (id, name, base_product_id, created_at, updated_at)
VALUES ('87ea063e-8d3d-4d5c-bfbd-9f58d2f0ebdb', 'lima', '28fb96f1-f213-4b93-9379-493b298c879b', NOW(), NOW());

-- Inventory
INSERT INTO
inventories (id, user_id, product_id, quantity, unit_of_measure_id, created_at, updated_at)
VALUES ('5e6da197-7762-448d-8e90-471fcc2c457d', '6bb774bd-fc80-40a9-a063-c7838209ec54', '87ea063e-8d3d-4d5c-bfbd-9f58d2f0ebdb', 10, 'd01637a9-d6a4-44c9-92dc-12f25363cbe9', NOW(), NOW());