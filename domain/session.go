package domain

import (
	"time"
)

// Session stores the session info of user
type Session struct {
	SessionID string    `db:"session_id"`
	UserID    int64     `db:"user_id"`
	LoginTime time.Time `db:"login_time"`
	Status    int32     `db:"status"`
}

type SessionRepository interface {
	Save(tenant string, session Session) error
}
