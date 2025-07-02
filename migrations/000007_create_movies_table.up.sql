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