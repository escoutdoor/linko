package entity

import "time"

type Product struct {
	ID         string
	CategoryID string

	Name        string
	Description string
	ImageURL    string

	PriceCents  int
	WeightGrams int

	IsAvailable bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
