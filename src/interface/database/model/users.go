package model

import "time"

type Users struct {
	UserId      string
	Email       string
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
