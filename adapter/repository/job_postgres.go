package repository

import (
	"context"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/pkg/errors"
)

type JobSQL struct {
	db SQL
}

func NewJobSQL(db SQL) JobSQL {
	return JobSQL{
		db: db,
	}
}

func (j JobSQL) Create(ctx context.Context, job domain.Job) (domain.Job, error) {
	var query = `
		INSERT INTO
			jobs (id, partner_id, title, status, category_id, expires_at, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
	`

	if err := j.db.ExecuteContext(
		ctx,
		query,
		job.ID(),
		job.PartnerID(),
		job.Title(),
		job.Status(),
		job.CategoryID(),
		job.ExpiresAt(),
		job.CreatedAt(),
	); err != nil {
		return domain.Job{}, errors.Wrap(err, "error creating job")
	}

	return job, nil
}
