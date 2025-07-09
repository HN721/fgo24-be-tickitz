CREATE TABLE transaction_detail (
    id SERIAL PRIMARY KEY,
    id_transaction INT REFERENCES transactions (id) ON DELETE CASCADE,
    costumer_name VARCHAR(100) NOT NULL,
    costumer_phone VARCHAR(20),
    seat VARCHAR(10) NOT NULL
);