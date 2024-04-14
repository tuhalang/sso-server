package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/tuhalang/authen/domain"
)

const pGDriverName = "pgx"

// PostgresStore is a postgres store
type PostgresStore struct {
	dbx *sqlx.DB
}

// GetByUsername retrieves user by username
func (store *PostgresStore) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := store.dbx.GetContext(ctx, &user, `SELECT ID, USERNAME, PASSWORD, STATUS FROM USERS WHERE USERNAME = $1`, username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, nil
	}

	return &user, nil
}

func (store *PostgresStore) SaveSession(session domain.Session) error {
	insertQuery := `INSERT INTO SESSIONS (USER_ID, SESSION_ID, LOGIN_TIME ,STATUS) VALUES ($1, $2, $3, $4)`
	store.dbx.MustExec(insertQuery, session.UserID, session.SessionID, session.LoginTime, session.Status)
	return nil
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
