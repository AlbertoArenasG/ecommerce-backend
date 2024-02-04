-- init.sql

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_url TEXT,
    description TEXT
);

ALTER TABLE products
ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT false;
