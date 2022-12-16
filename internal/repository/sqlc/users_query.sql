-- name: UserCreate :one
INSERT INTO users (
    login, password
) VALUES (
    $1, $2
) RETURNING id;
