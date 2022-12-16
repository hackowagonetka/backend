CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    login varchar NOT NULL UNIQUE,
    password varchar NOT NULL
);

CREATE TABLE points ( 
    id BIGSERIAL PRIMARY KEY,
    lat decimal, 
    lon decimal
);