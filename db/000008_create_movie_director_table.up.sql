CREATE TABLE movie_director (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies (id) ON DELETE CASCADE,
    director_id INT REFERENCES directors (id) ON DELETE CASCADE
);