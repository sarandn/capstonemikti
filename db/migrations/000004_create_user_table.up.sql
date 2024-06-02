CREATE TABLE users (
    user_id serial PRIMARY KEY,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    full_name varchar(255) NOT NULL,
    phone_num varchar(15) NOT NULL,
    address varchar(255) NOT NULL,
    role_id_fk int NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_roles FOREIGN KEY(role_id_fk) REFERENCES roles(role_id)
);