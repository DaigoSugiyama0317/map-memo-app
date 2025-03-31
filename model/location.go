package model

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID    uint    `json:"user_id"`
	Name      string  `gorm:"size:255;not null" json:"name"`
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
	Memo      string  `gorm:"type:text" json:"memo"`
	ImagePath string  `gorm:"size:512" json:"image_path"`

	Tags []Tag `gorm:"many2many:location_tags;constraint:OnDelete:CASCADE;" json:"tags,omitempty"`
}
