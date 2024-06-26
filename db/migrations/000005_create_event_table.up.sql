CREATE TABLE events (
    event_id serial PRIMARY KEY,
    user_id_fk int NOT NULL,
    event_name varchar(255) NOT NULL,
    image text,
    location varchar(255),
    longitude double precision,
    latitude double precision,
    date_start timestamp,
    date_end timestamp,
    price int,
    quantity_id_fk int,
    category_id_fk int,
    total_like int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id_fk) REFERENCES users(user_id),
    CONSTRAINT fk_quantity FOREIGN KEY(quantity_id_fk) REFERENCES quantities(quantity_id),
    CONSTRAINT fk_category FOREIGN KEY(category_id_fk) REFERENCES category(category_id)
);
