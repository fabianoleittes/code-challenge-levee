package output

import (
	"github.com/fabianoleittes/code-challenge-levee/domain"
)

// Job represents the output attributes
type Job struct {
	ID         string `json:"id"`
	PartnerID  string `json:"partner_id"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	CategoryID string `json:"category_id"`
	ExpiresAt  string `json:"expires_at"`
	CreatedAt  string `json:"created_at"`
}

// JobPresenter represents the interface contract
type JobPresenter interface {
	Output(domain.Job) Job
}
