package repository

import (
	"context"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type jobBSON struct {
	ID         string    `bson:"id"`
	PartnerID  string    `bson:"partner_id"`
	Title      string    `bson:"title"`
	Status     string    `bson:"status"`
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
		Status:     job.Status(),
		CategoryID: job.CategoryID(),
		ExpiresAt:  job.ExpiresAt(),
		CreatedAt:  job.CreatedAt(),
	}

	if err := j.db.Store(ctx, j.collectionName, jobBSON); err != nil {
		return domain.Job{}, errors.Wrap(err, "error creating job")
	}

	return job, nil
}

func (j JobNoSQL) FindAll(ctx context.Context) ([]domain.Job, error) {
	var jobsBSON = make([]jobBSON, 0)

	if err := j.db.FindAll(ctx, j.collectionName, bson.M{}, &jobsBSON); err != nil {
		switch err {
		case mongo.ErrNilDocument:
			return []domain.Job{}, errors.Wrap(domain.ErrJobNotFound, "error listing jobs")
		default:
			return []domain.Job{}, errors.Wrap(err, "error listing jobs")
		}
	}

	var jobs = make([]domain.Job, 0)

	for _, jobBSON := range jobsBSON {
		var job = domain.NewJob(
			domain.JobID(jobBSON.ID),
			jobBSON.PartnerID,
			jobBSON.Title,
			jobBSON.Status,
			jobBSON.CategoryID,
			jobBSON.ExpiresAt,
			jobBSON.CreatedAt,
		)

		jobs = append(jobs, job)
	}

	return jobs, nil
}
