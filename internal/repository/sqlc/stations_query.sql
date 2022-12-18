-- name: StationCreate :one
INSERT INTO stations (
    created_at, name, geoname, lon, lat, ref_user_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id;

-- name: StationGetList :many
SELECT * FROM stations WHERE ref_user_id = $1;

-- name: StationGetListByID :many
SELECT * FROM stations WHERE id = ANY($1::bigint[]);