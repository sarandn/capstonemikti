CREATE TABLE Payment (
    payment_id SERIAL PRIMARY KEY,
    order_id_fk INT,
    payment_date TIMESTAMP,
    amount_paid INT,
    payment_method VARCHAR(255),
    payment_status VARCHAR(255),
    CONSTRAINT fk_orders FOREIGN KEY (order_id_fk) REFERENCES orders(order_id)
);
