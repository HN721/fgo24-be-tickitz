CREATE TABLE movie_director (
    id VARCHAR(36) PRIMARY KEY,
    movie_id VARCHAR(36) REFERENCES movies (id) ON DELETE CASCADE,
    director_id VARCHAR(36) REFERENCES directors (id) ON DELETE CASCADE
);