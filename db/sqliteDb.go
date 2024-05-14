package db

import (
	"context"
	"database/sql"
	"time"
)

type SqliteDb struct {
	db *sql.DB
}

func (sql *SqliteDb) Set(ctx context.Context, key, value string, timeout time.Duration) error {
	// just an example
	return nil
}

func (sql *SqliteDb) Get(ctx context.Context, key string) (string, error) {
	// just an example
	return "", nil
}

func (sql *SqliteDb) Incr(ctx context.Context, key string) error {
	// just an example
	return nil
}

func NewSqliteDb(addr, password string, db int) (*SqliteDb, error) {
	// just an example
	return &SqliteDb{}, nil
}
