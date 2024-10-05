package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"index:idx_users_username_active,unique,where:deleted_at is null" json:"username"`
	Email    string `gorm:"index:idx_users_email_active,unique,where:deleted_at is null" json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8"`
}
