package model

import "time"

type Tag struct {
	ID        uint
	Name      string
	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
