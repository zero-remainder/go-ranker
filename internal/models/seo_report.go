package models

import (
	"gorm.io/gorm"
	"time"
)

type SEOReport struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	URL         string `gorm:"not null"`
	Title       string
	Description string
	Headings    string // JSON string or custom struct
	MetaTags    string // JSON string or custom struct
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
