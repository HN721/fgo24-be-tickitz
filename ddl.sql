-- Active: 1750554601825@@127.0.0.1@5434@movxtar

CREATE DATABASE movxtar;

CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20),
    profile_image VARCHAR(255),
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel sesi pengguna
CREATE TABLE sessions (
    id VARCHAR(36) PRIMARY KEY,
    token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_users VARCHAR(36) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE genres (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE directors (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE actors (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE movies (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    synopsis TEXT,
    background VARCHAR(255),
    poster VARCHAR(255),
    release_date TIMESTAMP,
    duration INTEGER CHECK (duration > 0),
    price INTEGER CHECK (price >= 0)
);

CREATE TABLE movie_genre (
    id VARCHAR(36) PRIMARY KEY,
    movie_id VARCHAR(36) REFERENCES movies (id) ON DELETE CASCADE,
    genre_id VARCHAR(36) REFERENCES genres (id) ON DELETE CASCADE
);

CREATE TABLE movie_director (
    id VARCHAR(36) PRIMARY KEY,
    movie_id VARCHAR(36) REFERENCES movies (id) ON DELETE CASCADE,
    director_id VARCHAR(36) REFERENCES directors (id) ON DELETE CASCADE
);

CREATE TABLE movie_actors (
    id VARCHAR(36) PRIMARY KEY,
    movie_id VARCHAR(36) REFERENCES movies (id) ON DELETE CASCADE,
    actor_id VARCHAR(36) REFERENCES actors (id) ON DELETE CASCADE
);

CREATE TABLE cinema (
    id VARCHAR(39) PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE payment_method (
    id VARCHAR(39) PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE transactions (
    id VARCHAR(36) PRIMARY KEY,
    time TIMESTAMP NOT NULL,
    date TIMESTAMP NOT NULL,
    cinema VARCHAR(255) NOT NULL,
    price_total INTEGER CHECK (price_total >= 0),
    user_id VARCHAR(36) REFERENCES users (id) ON DELETE SET NULL,
    movie_id VARCHAR(36) REFERENCES movies (id) ON DELETE SET NULL,
    id_cinema VARCHAR(39),
    id_payment_method VARCHAR(39),
    CONSTRAINT fk_cinema FOREIGN KEY (id_cinema) REFERENCES cinema (id),
    CONSTRAINT fk_payment_method FOREIGN KEY (id_payment_method) REFERENCES payment_method (id)
);

CREATE TABLE transaction_detail (
    id VARCHAR(36) PRIMARY KEY,
    id_transaction VARCHAR(36) REFERENCES transactions (id) ON DELETE CASCADE,
    costumer_name VARCHAR(100) NOT NULL,
    costumer_phone VARCHAR(20),
    seat VARCHAR(10) NOT NULL
);

CREATE TABLE history_transaction (
    id VARCHAR(36) PRIMARY KEY,
    transaction_id VARCHAR(36) REFERENCES transactions (id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL,
    updated_by VARCHAR(50),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    note TEXT
);