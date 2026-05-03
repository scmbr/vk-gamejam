CREATE TABLE IF NOT EXISTS child_profiles (
    id SERIAL PRIMARY KEY,

    user_id INT UNIQUE NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    child_name TEXT,
    child_gender TEXT,
    parent_pin TEXT,
    has_pet BOOLEAN DEFAULT FALSE,

    is_first_launch BOOLEAN DEFAULT TRUE,
    last_login TIMESTAMP,
    last_logout TIMESTAMP
);