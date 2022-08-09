CREATE TABLE users (
    id UUID,
	first_name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30) NOT NULL,
	address VARCHAR(100) NOT NULL,
	phone VARCHAR(14) NOT NULL,
	email VARCHAR(50) NOT NULL UNIQUE,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL,
	roles TEXT[] NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id)
);

CREATE TABLE units_of_measure (
    id UUID,
	name VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id)
);

CREATE TABLE base_products (
    id UUID,
	name VARCHAR(30) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id)
);

CREATE TABLE products (
    id UUID,
	name VARCHAR(30) NOT NULL UNIQUE,
    base_product_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id),
	UNIQUE(name, base_product_id),
	CONSTRAINT base_product_id FOREIGN KEY (base_product_id) REFERENCES base_products(id)
);

CREATE TABLE inventories (
    id UUID,
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
	quantity INTEGER NOT NULL,
    unit_of_measure_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id),
	UNIQUE(user_id, product_id, unit_of_measure_id),
	CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT product_id FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT unit_of_measure_id FOREIGN KEY (unit_of_measure_id) REFERENCES units_of_measure(id)
);
