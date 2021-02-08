package repository

import (
	"context"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/pkg/errors"
)

type jobBSON struct {
	ID         string    `bson:"id"`
	PartnerID  string    `bson:"partner_id"`
	Title      string    `bson:"title"`
	CategoryID string    `bson:"category_id"`
	ExpiresAt  time.Time `bson:"expires_at"`
	CreatedAt  time.Time `bson:"created_at"`
}

type JobNoSQL struct {
	collectionName string
	db             NoSQL
}

func NewJobNoSQL(db NoSQL) JobNoSQL {
	return JobNoSQL{
		db:             db,
		collectionName: "jobs",
	}
}

func (j JobNoSQL) Create(ctx context.Context, job domain.Job) (domain.Job, error) {
	var jobBSON = jobBSON{
		ID:         job.ID().String(),
		PartnerID:  job.PartnerID(),
		Title:      job.Title(),
		CategoryID: job.CategoryID(),
		ExpiresAt:  job.ExpiresAt(),
		CreatedAt:  job.CreatedAt(),
	}

	if err := j.db.Store(ctx, j.collectionName, jobBSON); err != nil {
		return domain.Job{}, errors.Wrap(err, "error creating job")
	}

	return job, nil
}
