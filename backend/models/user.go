package models

import "time"

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name,omitempty"`
	IsVerified  bool      `json:"is_verified"`
	LastLogin   time.Time `json:"last_login,omitempty"`
	IsActive    bool      `json:"is_active"`
}
