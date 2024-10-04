package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"size:100;not null"`
	Stock int     `gorm:"not null"`
	Price float64 `gorm:"not null"`
}
