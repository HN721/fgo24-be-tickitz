ALTER TABLE users ADD COLUMN profile_id INTEGER;

ALTER TABLE users
ADD CONSTRAINT fk_users_profiles FOREIGN KEY (profile_id) REFERENCES profiles (id) ON DELETE SET NULL;