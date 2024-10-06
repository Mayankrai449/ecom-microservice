package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint         `json:"user_id"`
	Products    ProductItems `gorm:"type:jsonb" json:"products"`
	TotalPrice  float64      `json:"total_price"`
	Paid        bool         `json:"paid"`
	OrderStatus string       `json:"order_status" validate:"required,oneof=Pending Completed pending completed"`
}

const (
	StatusPending   = "Pending"
	StatusCompleted = "Completed"
)
