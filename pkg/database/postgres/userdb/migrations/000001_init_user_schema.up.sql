CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar NOT NULL,
    email varchar NOT NULL,
    firstname varchar NOT NULL,
    lastname varchar NOT NULL,
    phone_no varchar NOT NULL,
    private_account boolean NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE friends (
    id SERIAL PRIMARY KEY,
    user_id1 INTEGER REFERENCES users(id),
    user_id2 INTEGER REFERENCES users(id),
    status VARCHAR DEFAULT 'Pending', -- Pending, Accepted, Declined, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE friends ADD CONSTRAINT unique_friendship UNIQUE (user_id1, user_id2);