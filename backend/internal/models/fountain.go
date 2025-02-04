package models

type Fountain struct {
	ID        int64   `json:"id"`
	State     string  `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
