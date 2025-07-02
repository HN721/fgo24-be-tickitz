CREATE TABLE transactions (
    id VARCHAR(36) PRIMARY KEY,
    time TIMESTAMP NOT NULL,
    date TIMESTAMP NOT NULL,
    price_total INTEGER CHECK (price_total >= 0),
    user_id VARCHAR(36) REFERENCES users (id) ON DELETE SET NULL,
    movie_id VARCHAR(36) REFERENCES movies (id) ON DELETE SET NULL,
    id_cinema VARCHAR(39),
    id_payment_method VARCHAR(39),
    CONSTRAINT fk_cinema FOREIGN KEY (id_cinema) REFERENCES cinema (id),
    CONSTRAINT fk_payment_method FOREIGN KEY (id_payment_method) REFERENCES payment_method (id)
);