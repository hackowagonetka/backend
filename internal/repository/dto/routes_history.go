package dto

import "github.com/cridenour/go-postgis"

type RoutesHistoryData struct {
	Points []postgis.Point `json:"points"`
}
