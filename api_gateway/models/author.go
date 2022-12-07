package models

import "time"

// Author ...
type Author struct {
	ID        string     `json:"id"`
	Fullname  string     `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateAuthorModel ...
type CreateAuthorModel struct {
	Fullname string `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
}
