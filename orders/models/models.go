package models

import "gorm.io/gorm"

type ProductItem struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type Order struct {
	gorm.Model
	UserID      uint          `json:"user_id"`
	Products    []ProductItem `gorm:"type:json" json:"products"`
	TotalPrice  float64       `json:"total_price"`
	Paid        bool          `json:"paid"`
	OrderStatus string        `json:"order_status" validate:"required,oneof=Pending Completed pending completed"`
}

const (
	StatusPending   = "Pending"
	StatusCompleted = "Completed"
)
