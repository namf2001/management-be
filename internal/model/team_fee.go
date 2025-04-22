package model

import (
	"time"
)

// TeamFee represents a team fee entity
type TeamFee struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
