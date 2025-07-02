DROP TABLE IF EXISTS profile CASCADE;

CREATE TABLE profile (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255),
    phone_number VARCHAR(20),
    profile_image VARCHAR(255),
    id_user INT REFERENCES users (id) ON DELETE CASCADE
)