-- name: UserCreate :one
INSERT INTO users (
    login, password
) VALUES (
    $1, $2
) RETURNING id;


-- name: UserGet :one
SELECT * FROM users WHERE login = $1;