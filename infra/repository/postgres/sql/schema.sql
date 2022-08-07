CREATE TABLE base_products (
    id UUID PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE inventories (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
	quantity INTEGER NOT NULL,
    unit_of_measure_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT product_id FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT unit_of_measure_id FOREIGN KEY (unit_of_measure_id) REFERENCES units_of_measure(id)
);

CREATE TABLE products (
    id UUID PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
    base_product_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT base_product_id FOREIGN KEY (base_product_id) REFERENCES base_products(id)
);

CREATE TABLE units_of_measure (
    id UUID PRIMARY KEY,
	name VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE customers (
	id UUID PRIMARY KEY,
	first_name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30) NOT NULL,
	address VARCHAR(100) NOT NULL,
	phone VARCHAR(14) NOT NULL,
	email VARCHAR(50) NOT NULL,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
	first_name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30) NOT NULL,
	address VARCHAR(100) NOT NULL,
	phone VARCHAR(14) NOT NULL,
	email VARCHAR(50) NOT NULL,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL,
	role_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE roles (
	id UUID PRIMARY KEY,
	name VARCHAR(15),
    created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);
