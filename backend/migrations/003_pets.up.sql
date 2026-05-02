CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,

    user_id INT UNIQUE NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    name TEXT,
    type TEXT,
    gender TEXT,

    level INT DEFAULT 1,
    xp FLOAT DEFAULT 0,

    knowledge FLOAT DEFAULT 0,
    energy FLOAT DEFAULT 0,
    creativity FLOAT DEFAULT 0,
    sport FLOAT DEFAULT 0,

    last_online TIMESTAMP NOT NULL DEFAULT NOW()
);