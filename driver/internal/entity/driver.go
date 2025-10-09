package entity

import "time"

type Driver struct {
	ID             string
	UserID         string
	TotalRatingSum float64
	ReviewCount    int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
