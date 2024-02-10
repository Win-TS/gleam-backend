CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    firstname VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL,
    phone_no VARCHAR NOT NULL,
    private_account BOOLEAN NOT NULL DEFAULT false,
    nationality VARCHAR NOT NULL,
    age INTEGER NOT NULL,
    birthday DATE NOT NULL,
    gender VARCHAR NOT NULL,
    photourl VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE friends (
    id SERIAL PRIMARY KEY,
    user_id1 INTEGER REFERENCES users(id),
    user_id2 INTEGER REFERENCES users(id),
    status VARCHAR DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE friends ADD CONSTRAINT unique_friendship UNIQUE (user_id1, user_id2);