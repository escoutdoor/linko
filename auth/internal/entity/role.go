package entity

import "time"

type Role struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRole struct {
	ID     string
	UserID string
	RoleID string
}
