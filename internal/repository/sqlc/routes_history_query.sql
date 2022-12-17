-- name: RoutesHistoryCreate :one
INSERT INTO routes_history ( created_at, data, ref_user_id ) VALUES ( $1, $2, $3 ) RETURNING id;
-- name: RoutesHistoryGet :many
SELECT  *
FROM routes_history
WHERE ref_user_id = ?
ORDER BY id DESC
LIMIT 10;
-- name: RoutesDistance :one
SELECT  round (CAST( ST_DistanceSphere( ST_MakePoint($1,$2),ST_MakePoint($3,$4)) AS numeric )) AS meters;