package database

import (
	"errors"

	"github.com/fabianoleittes/code-challenge-levee/adapter/repository"
)

var (
	errInvalidNoSQLDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceMongoDB int = iota
)

// NewDatabaseNoSQLFactory will create a new instance of MongoDB
func NewDatabaseNoSQLFactory(instance int) (repository.NoSQL, error) {
	switch instance {
	case InstanceMongoDB:
		return NewMongoHandler(newConfigMongoDB())
	default:
		return nil, errInvalidNoSQLDatabaseInstance
	}
}
