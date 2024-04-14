package repository

import (
	"errors"
	"github.com/tuhalang/authen/domain"
	"github.com/tuhalang/authen/store"
)

type sessionRepository struct {
	databases map[string]store.Store
}

// NewSessionRepository inits the user repository
func NewSessionRepository(dbs map[string]store.Store) domain.SessionRepository {
	return &sessionRepository{
		databases: dbs,
	}
}

func (repository *sessionRepository) Save(tenant string, session domain.Session) error {
	database, ok := repository.databases[tenant]
	if !ok {
		return errors.New("Unknown space: " + tenant)
	}
	return database.SaveSession(session)
}
