package domain

import (
	"context"
)

// User stores the user's data
type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Status   int32  `db:"status"`
}

type UserRepository interface {
	GetByUsername(ctx context.Context, tenant, username string) (*User, error)
}
