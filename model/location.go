package model

import "time"

type Location struct {
	ID        uint
	Name string
	Lat int
	Ing int
	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}