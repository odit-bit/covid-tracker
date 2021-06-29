package model

import "time"

type CovidData struct {
	// ID        string    `json:"id"`
	Name      string    `json:"name"`
	Total     int       `json:"total"`
	Inap      int       `json:"inap"`
	Mandiri   int       `json:"mandiri"`
	Sembuh    int       `json:"sembuh"`
	Meninggal int       `json:"meninggal"`
	UpdatedAt time.Time `json:"updatedAt"`
}
