package presenter

import (
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/fabianoleittes/code-challenge-levee/usecase"
)

type findAllJobPresenter struct{}

func NewFindAllJobPresenter() usecase.FindAllJobPresenter {
	return findAllJobPresenter{}
}

func (a findAllJobPresenter) Output(jobs []domain.Job) []usecase.FindAllJobOutput {
	var o = make([]usecase.FindAllJobOutput, 0)

	for _, job := range jobs {
		o = append(o, usecase.FindAllJobOutput{
			ID:         job.ID().String(),
			PartnerID:  job.PartnerID(),
			Title:      job.Title(),
			Status:     job.Status(),
			CategoryID: job.CategoryID(),
			ExpiresAt:  job.ExpiresAt().Format(time.RFC3339),
			CreatedAt:  job.CreatedAt().Format(time.RFC3339),
		})
	}

	return o
}
