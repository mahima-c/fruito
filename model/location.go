package model

type Location struct {
	LocationID  int     `json:"location_id"`
	Name        string  `json:"name"` // e.g., "Main Building", "Emergency Wing", "Outpatient Clinic"
	Address     string  `json:"address"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	ZipCode     string  `json:"zip_code"`
	Country     string  `json:"country"`
	Latitude    float64 `json:"latitude,omitempty"`  // Optional latitude
	Longitude   float64 `json:"longitude,omitempty"` // Optional longitude
	PhoneNumber string  `json:"phone_number,omitempty"`
	Notes       string  `json:"notes,omitempty"`
}
