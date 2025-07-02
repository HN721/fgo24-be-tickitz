CREATE TABLE history_transaction (
    id VARCHAR(36) PRIMARY KEY,
    transaction_id VARCHAR(36) REFERENCES transactions (id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL,
    updated_by VARCHAR(50),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    note TEXT
);