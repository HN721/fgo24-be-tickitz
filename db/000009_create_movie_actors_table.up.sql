CREATE TABLE movie_actors (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies (id) ON DELETE CASCADE,
    actor_id INT REFERENCES actors (id) ON DELETE CASCADE
);