package entity

import "time"

type Driver struct {
	ID          string
	UserID      string
	Rating      float32
	ReviewCount int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
