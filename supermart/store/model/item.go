package model

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	SuperMartID string  `json:"super_mart_id"`
}
