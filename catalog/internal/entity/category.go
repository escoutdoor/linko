package entity

import "time"

type Category struct {
	ID      string
	StoreID string

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}
