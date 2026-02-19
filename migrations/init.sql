CREATE TABLE users (id SERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE);
CREATE TABLE carts (user_id INT REFERENCES users(id), product_id INT, quantity INT);
CREATE TABLE products (id SERIAL PRIMARY KEY, name VARCHAR(100), stock INT);
CREATE TABLE orders (id SERIAL PRIMARY KEY, user_id INT REFERENCES users(id), product_id INT, quantity INT, status VARCHAR(20));