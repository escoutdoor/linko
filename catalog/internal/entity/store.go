package entity

import "time"

type Store struct {
	ID string

	Name        string
	Description string
	ImageURL    string
	Address     string

	IsActive bool

	Latitude  float64
	Longitude float64

	CreatedAt time.Time
	UpdatedAt time.Time
}
