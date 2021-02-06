package input

import (
	"time"
)

// Job represents the input attributes
type Job struct {
	PartnerID  string    `json:"partner_id" validate:"required"`
	Title      string    `json:"title" validate:"required"`
	Status     string    `json:"status"`
	CategoryID string    `json:"category_id" validate:"required"`
	ExpiresAt  time.Time `json:"expires_at" validate:"required"`
}
