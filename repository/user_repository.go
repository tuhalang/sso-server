package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/tuhalang/authen/domain"
	"github.com/tuhalang/authen/store"
)

type userRepository struct {
	databases map[string]store.Store
}

// NewUserRepository inits the user repository
func NewUserRepository(dbs map[string]store.Store) domain.UserRepository {
	return &userRepository{
		databases: dbs,
	}
}

func (repository *userRepository) GetByUsername(ctx context.Context, tenant, username string) (*domain.User, error) {
	log := zerolog.Ctx(ctx)
	log.Info().Msgf("Handle login for user[%s] in space[%s]", username, tenant)
	database, ok := repository.databases[tenant]
	if !ok {
		return nil, errors.New("Unknown space: " + tenant)
	}
	return database.GetByUsername(ctx, username)
}
