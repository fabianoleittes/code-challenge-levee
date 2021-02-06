package usecase

import (
	"context"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/fabianoleittes/code-challenge-levee/usecase/input"
	"github.com/fabianoleittes/code-challenge-levee/usecase/output"
)

// CreateJob represents the interface contract
type CreateJob interface {
	Execute(context.Context, input.Job) (output.Job, error)
}

// CreateJobInteractor represents the interactor contract
type CreateJobInteractor struct {
	repo       domain.JobRepository
	presenter  output.JobPresenter
	ctxTimeout time.Duration
}

// NewCreateJobInteractor will create a new instance of CreateJobInteractor
func NewCreateJobInteractor(
	repo domain.JobRepository,
	presenter output.JobPresenter,
	t time.Duration,
) CreateJobInteractor {
	return CreateJobInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute will create a new Job by input params
func (j CreateJobInteractor) Execute(ctx context.Context, input input.Job) (output.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, j.ctxTimeout)
	defer cancel()

	var job = domain.NewJob(
		domain.JobID(domain.NewUUID()),
		input.PartnerID,
		input.Title,
		"draft",
		input.CategoryID,
		input.ExpiresAt,
		time.Now(),
	)

	job, err := j.repo.Create(ctx, job)

	if err != nil {
		return j.presenter.Output(domain.Job{}), err
	}

	return j.presenter.Output(job), nil
}
