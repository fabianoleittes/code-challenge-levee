package domain

import "github.com/google/uuid"

// ID domain ID
type ID = uuid.UUID

//NewID create a new domain ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID convert a string to an domain ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
