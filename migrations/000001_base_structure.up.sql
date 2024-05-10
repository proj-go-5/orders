CREATE TYPE status_type AS ENUM ('not_active', 'active', 'updated', 'finished');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    status status_type NOT NULL,
    total_price INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_history (
    order_id INTEGER NOT NULL,
    status status_type NOT NULL,
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