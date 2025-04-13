package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID uint   `json:"user_id"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Color  string `gorm:"size:50" json:"color"`
	// Index  uint   `json:"index"`

	Locations []Location `gorm:"many2many:location_tags;constraint:OnDelete:CASCADE;" json:"locations,omitempty"`
}

type TagResponse struct {
	ID    uint   `json:id`
	Name  string `json:name`
	Color string `json:color`
}
