CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL, 
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT NOT NULL,
    product_images TEXT[], 
    compressed_product_images TEXT[],
    product_price NUMERIC(10, 2) NOT NULL CHECK (product_price > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

