CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    text_comment TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id),
    corrected_ai BOOLEAN DEFAULT FALSE
);
