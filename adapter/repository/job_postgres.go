package repository

import (
	"context"
	"time"

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

func (j JobSQL) FindAll(ctx context.Context) ([]domain.Job, error) {
	var query = "SELECT * FROM jobs"

	rows, err := j.db.QueryContext(ctx, query)
	if err != nil {
		return []domain.Job{}, errors.Wrap(err, "error listing jobs")
	}

	var jobs = make([]domain.Job, 0)
	for rows.Next() {
		var (
			ID         string
			PartnerID  string
			Title      string
			Status     string
			CategoryID string
			ExpiresAt  time.Time
			createdAt  time.Time
		)

		if err = rows.Scan(&ID, &PartnerID, &Title, &Status, &CategoryID, &ExpiresAt, &createdAt); err != nil {
			return []domain.Job{}, errors.Wrap(err, "error listing jobs")
		}

		jobs = append(jobs, domain.NewJob(
			domain.JobID(ID),
			PartnerID,
			Title,
			Status,
			CategoryID,
			ExpiresAt,
			createdAt,
		))
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return []domain.Job{}, err
	}

	return jobs, nil
}
