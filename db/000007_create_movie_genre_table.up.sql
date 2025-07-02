CREATE TABLE movie_genre (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies (id) ON DELETE CASCADE,
    genre_id INT REFERENCES genres (id) ON DELETE CASCADE
);