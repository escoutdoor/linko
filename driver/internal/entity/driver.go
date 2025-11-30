package entity

import "time"

type Driver struct {
	ID     string
	UserID string

	Status int32

	Rating      float32
	ReviewCount int32

	Vehicle Vehicle

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Vehicle struct {
	Type        int32
	Model       string
	PlateNumber string
	Color       string
}
