CREATE TABLE history_transaction (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions (id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL,
    note TEXT updated_by VARCHAR(50),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);