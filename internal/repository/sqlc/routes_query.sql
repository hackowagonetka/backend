
-- name: RouteDistance :one
SELECT  round (CAST( ST_DistanceSphere( ST_MakePoint($1,$2),ST_MakePoint($3,$4)) AS numeric )) AS meters;