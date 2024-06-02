CREATE TABLE roles (
    role_id serial PRIMARY KEY,
    role_name varchar(50) CHECK (role_name IN ('penjual', 'pembeli')) NOT NULL
);