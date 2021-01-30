// Package domain provides encapsulate business rules
package domain

import (
	"time"
)

type job struct {
	ID         ID
	PartnerId  ID
	Title      string
	categoryId ID
	ExpiresAt  time.Time
}
