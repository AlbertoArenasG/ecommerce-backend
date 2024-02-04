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

CREATE TABLE IF NOT EXISTS shopping_carts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS shopping_carts_items (
    cart_id INT,
    product_id INT,
    quantity INT DEFAULT 1,
    is_deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (cart_id) REFERENCES shopping_carts(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    PRIMARY KEY (cart_id, product_id)
);
