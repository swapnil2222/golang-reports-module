package models

import "time"

// Report model
type Report struct {
	ID           string    `json:"_id"`
	Product      string    `json:"product"`
	Manufacturer string    `json:"manufacturer"`
	Category     string    `json:"category"`
	VideoTitle   string    `json:"videoTitle"`
	VideoCode    string    `json:"videoCode"`
	DateReleased time.Time `json:"dateReleased"`
	Rating       float64   `json:"rating"`
	V            int       `json:"__v"`
}
