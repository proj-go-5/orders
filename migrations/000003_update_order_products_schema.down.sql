CREATE TABLE order_products_new (
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

INSERT INTO order_products_new (order_id, product_id, price, quantity)
SELECT order_id, product_id, price, quantity
FROM order_products;

DROP TABLE order_products;

ALTER TABLE order_products_new RENAME TO order_products;