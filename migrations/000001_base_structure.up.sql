CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    status INTEGER NOT NULL,
    total_price INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_history (
    order_id INTEGER NOT NULL,
    status INTEGER NOT NULL,
    comment VARCHAR(128),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE order_products (
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);