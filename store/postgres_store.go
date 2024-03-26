package store

import (
	"context"
	"github.com/jmoiron/sqlx"
)

const pGDriverName = "pgx"

// PostgresStore is a postgres store
type PostgresStore struct {
	dbx *sqlx.DB
}

// NewPostgresStore create a postgres store
func NewPostgresStore(ctx context.Context, dbURL string) (Store, error) {
	dbx, err := sqlx.ConnectContext(ctx, pGDriverName, dbURL)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{dbx: dbx}, nil
}

// Close postgres connection
func (store *PostgresStore) Close() error {
	return store.dbx.Close()
}
