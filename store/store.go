package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/tuhalang/authen/config"
)

type Store interface {
	Close() error
}

func NewStore(ctx context.Context, dbConfig config.DatabaseConfig) (Store, error) {
	switch dbConfig.DBDriver {
	case pGDriverName:
		dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
		return NewPostgresStore(ctx, dbURL)
	default:
		return nil, errors.New("Unknown database driver " + dbConfig.DBDriver)
	}

}
