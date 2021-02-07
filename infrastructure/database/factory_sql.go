package database

import (
	"errors"

	"github.com/fabianoleittes/code-challenge-levee/adapter/repository"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

const (
	InstancePostgres int = iota
)

// NewDatabaseSQLFactory will create a new instance of PostgreSQL
func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstancePostgres:
		return NewPostgresHandler(newConfigPostgres())
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
