ALTER TABLE ORDERS DROP COLUMN customer_id;

ALTER TABLE ORDERS ADD customer_info jsonb NOT NULL;    