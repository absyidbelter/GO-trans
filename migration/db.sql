DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO postgres;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    number VARCHAR(255) NOT NULL UNIQUE,
    balance INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

DROP TABLE wallets;
-- Membuat tabel Transaction
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount INT NOT NULL,
    destination_id VARCHAR(255) NOT NULL,
    history TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    payment_method_type VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (destination_id) REFERENCES wallets(number)
);

-- Insert 3 user ke dalam tabel users
INSERT INTO users (name, email, password)
VALUES ('bima', 'bima@example.com', 'password123'),
       ('budiatun', 'budi@example.com', 'mypassword'),
       ('markzus', 'markzus@example.com', 'securepassword');

INSERT INTO wallets (user_id, number, balance)
VALUES
    (1, '1234567890', 1000),
    (2, '0987654321', 500),
    (3, '9876543210', 200)



