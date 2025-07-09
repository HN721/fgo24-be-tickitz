CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP NOT NULL,
    date TIMESTAMP NOT NULL,
    price_total INTEGER CHECK (price_total >= 0),
    user_id INT REFERENCES users (id) ON DELETE SET NULL,
    movie_id INT REFERENCES movies (id) ON DELETE SET NULL,
    id_cinema INT,
    id_payment_method INT,
    CONSTRAINT fk_cinema FOREIGN KEY (id_cinema) REFERENCES cinema (id),
    CONSTRAINT fk_payment_method FOREIGN KEY (id_payment_method) REFERENCES payment_method (id)
);