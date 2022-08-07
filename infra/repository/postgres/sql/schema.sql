CREATE TABLE base_products (
    id UUID,
	name VARCHAR(30) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id)
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
	CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT product_id FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT unit_of_measure_id FOREIGN KEY (unit_of_measure_id) REFERENCES units_of_measure(id)
);

CREATE TABLE products (
    id UUID,
	name VARCHAR(30) UNIQUE NOT NULL,
    base_product_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id),
	CONSTRAINT base_product_id FOREIGN KEY (base_product_id) REFERENCES base_products(id)
);

CREATE TABLE units_of_measure (
    id UUID,
	name VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id)
);

CREATE TABLE users (
    id UUID,
	first_name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30) NOT NULL,
	address VARCHAR(100) NOT NULL,
	phone VARCHAR(14) NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL,
	roles TEXT[] NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY (id),
	CONSTRAINT role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);
