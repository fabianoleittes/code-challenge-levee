package presenter

import (
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/fabianoleittes/code-challenge-levee/usecase/output"
)

type jobPresenter struct{}

func NewJobPresenter() jobPresenter {
	return jobPresenter{}
}

func (j jobPresenter) Output(job domain.Job) output.Job {
	return output.Job{
		ID:         job.ID().String(),
		PartnerID:  job.PartnerID(),
		Title:      job.Title(),
		Status:     job.Status(),
		CategoryID: job.CategoryID(),
		ExpiresAt:  job.ExpiresAt().Format(time.RFC3339),
		CreatedAt:  job.CreatedAt().Format(time.RFC3339),
	}
}
