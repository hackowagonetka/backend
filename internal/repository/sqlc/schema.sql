CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    login varchar NOT NULL UNIQUE,
    password varchar NOT NULL
);

CREATE TABLE stations (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    name varchar NOT NULL,
    geoname varchar NOT NULL,
    lon float NOT NULL,
    lat float NOT NULL,
    ref_user_id bigint REFERENCES users NOT NULL
);