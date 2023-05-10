-- File: db/migrations/20230430000001_create_users_table.up.sql
CREATE TABLE products (
                       id VARCHAR(255) PRIMARY KEY,
                       created_at TIMESTAMP NOT NULL,
                       description VARCHAR(255) NOT NULL,
                       price_amount INTEGER NOT NULL,
                       price_currency VARCHAR(255) NOT NULL
);
