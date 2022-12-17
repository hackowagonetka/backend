CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    login varchar NOT NULL UNIQUE,
    password varchar NOT NULL
);

CREATE TABLE routes_history ( 
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    data jsonb,
    ref_user_id integer REFERENCES users NOT NULL
);