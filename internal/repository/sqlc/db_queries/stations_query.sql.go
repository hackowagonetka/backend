// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: stations_query.sql

package db_queries

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const stationCreate = `-- name: StationCreate :one
INSERT INTO stations (
    created_at, name, geoname, lon, lat, ref_user_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id
`

type StationCreateParams struct {
	CreatedAt time.Time
	Name      string
	Geoname   string
	Lon       float64
	Lat       float64
	RefUserID int64
}

func (q *Queries) StationCreate(ctx context.Context, arg StationCreateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, stationCreate,
		arg.CreatedAt,
		arg.Name,
		arg.Geoname,
		arg.Lon,
		arg.Lat,
		arg.RefUserID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const stationGetList = `-- name: StationGetList :many
SELECT id, created_at, name, geoname, lon, lat, ref_user_id FROM stations WHERE ref_user_id = $1
`

func (q *Queries) StationGetList(ctx context.Context, refUserID int64) ([]Station, error) {
	rows, err := q.db.QueryContext(ctx, stationGetList, refUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Station
	for rows.Next() {
		var i Station
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Geoname,
			&i.Lon,
			&i.Lat,
			&i.RefUserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const stationGetListByID = `-- name: StationGetListByID :many
SELECT id, created_at, name, geoname, lon, lat, ref_user_id FROM stations WHERE id = ANY($1::bigint[])
`

func (q *Queries) StationGetListByID(ctx context.Context, dollar_1 []int64) ([]Station, error) {
	rows, err := q.db.QueryContext(ctx, stationGetListByID, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Station
	for rows.Next() {
		var i Station
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Geoname,
			&i.Lon,
			&i.Lat,
			&i.RefUserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}