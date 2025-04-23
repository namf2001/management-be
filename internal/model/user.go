package model

import "time"

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	FullName  string     `json:"full_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
