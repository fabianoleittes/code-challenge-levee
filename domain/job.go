// Package domain provides encapsulate business rules
package domain

import (
	"context"
	"time"
)

// JobID type
type JobID string

func (j JobID) String() string {
	return string(j)
}

// Job represents the job entity's attributes
type Job struct {
	id         JobID
	partnerID  string
	title      string
	categoryID string
	status     string
	expiresAt  time.Time
	createdAt  time.Time
}

// JobRepository represents the interface contract
type JobRepository interface {
	Create(context.Context, Job) (Job, error)
}

//NewJob create a new Job
func NewJob(ID JobID, partnerID string, title string, status string, categoryID string,
	expiresAt time.Time, createdAt time.Time) Job {
	return Job{
		id:         ID,
		partnerID:  partnerID,
		title:      title,
		categoryID: categoryID,
		status:     status,
		expiresAt:  expiresAt,
		createdAt:  createdAt,
	}
}

// ID returns id
func (j Job) ID() JobID {
	return j.id
}

// PartnerID returns partnerID
func (j Job) PartnerID() string {
	return j.partnerID
}

// Title returns title
func (j Job) Title() string {
	return j.title
}

// CategoryID returns categoryID
func (j Job) CategoryID() string {
	return j.categoryID
}

// Status returns status
func (j Job) Status() string {
	return j.status
}

// ExpiresAt returns expiresAt
func (j Job) ExpiresAt() time.Time {
	return j.expiresAt
}

// CreatedAt returns createdAt
func (j Job) CreatedAt() time.Time {
	return j.createdAt
}
