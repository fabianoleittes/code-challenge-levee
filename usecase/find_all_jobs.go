package usecase

import (
	"context"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
)

type (
	// Input port
	FindAllJobUseCase interface {
		Execute(context.Context) ([]FindAllJobOutput, error)
	}

	// Output port
	FindAllJobPresenter interface {
		Output([]domain.Job) []FindAllJobOutput
	}

	// OutputData
	FindAllJobOutput struct {
		ID         string `json:"id"`
		PartnerID  string `json:"partner_id"`
		Title      string `json:"title"`
		Status     string `json:"status"`
		CategoryID string `json:"category_id"`
		ExpiresAt  string `json:"expires_at"`
		CreatedAt  string `json:"created_at"`
	}

	findAllJobInteractor struct {
		repo       domain.JobRepository
		presenter  FindAllJobPresenter
		ctxTimeout time.Duration
	}
)

// NewFindAllJobInteractor creates new findAllJobInteractor with its dependencies
func NewFindAllJobInteractor(
	repo domain.JobRepository,
	presenter FindAllJobPresenter,
	t time.Duration,
) FindAllJobUseCase {
	return findAllJobInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a findAllJobInteractor) Execute(ctx context.Context) ([]FindAllJobOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	jobs, err := a.repo.FindAll(ctx)
	if err != nil {
		return a.presenter.Output([]domain.Job{}), err
	}

	return a.presenter.Output(jobs), nil
}
