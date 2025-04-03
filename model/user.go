package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Username string `gorm:"size:255;not null;uniqueIndex" json:"username"`
	Email    string `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`

	Locations []Location `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"locations,omitempty"`
	Tags      []Tag      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tags,omitempty"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
