CREATE TABLE orders (
    order_id serial PRIMARY KEY,
    user_id_fk int NOT NULL,
    order_date timestamp DEFAULT CURRENT_TIMESTAMP,
    total_amount int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users_order FOREIGN KEY(user_id_fk) REFERENCES users(user_id)
);