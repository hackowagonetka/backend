package repository_dto

type RoutesHistoryPoint struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Lon  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

type RoutesHistoryData []RoutesHistoryPoint
