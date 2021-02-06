package domain

import "github.com/google/uuid"

// NewUUID returns a new uuid
func NewUUID() string {
	return uuid.New().String()
}
