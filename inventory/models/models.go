package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"size:100;not null" json:"name"`
	Stock int     `gorm:"not null" json:"stock"`
	Price float64 `gorm:"not null" json:"price"`
}
