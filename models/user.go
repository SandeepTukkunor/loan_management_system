package models

import (
	"github.com/google/uuid"
)

type Users struct {
	UserID         uuid.UUID `json:"user_id"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Mobile         int       `json:"mobile"`
	IsActive       bool      `json:"is_active"`
	IsStaff        bool      `json:"is_staff"`
	ISEmailVerified  bool      `json:"email_verified"`
	IsMobileVerified bool      `json:"mobile_verified"`
	InfoID         uuid.UUID `json:"info_id"`
}
