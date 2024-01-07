package models

type Car struct {
	ID       string `json:"id"`
	Model    string `json:"model"`
	Brand    string `json:"brand"`
	Year     int    `json:"year"`
	DriverID string `json:"driver_id"`
}
