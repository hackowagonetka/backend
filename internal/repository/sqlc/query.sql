-- name: PointCreate :one
INSERT INTO points (
    lat, lon
) VALUES (
    $1, $2
) 
RETURNING *;