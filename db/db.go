package db

import (
	"context"
	"time"
)

type Db interface {
	Set(ctx context.Context, key, value string, timeout time.Duration) error

	Get(ctx context.Context, key string) (string, error)

	Incr(ctx context.Context, key string) error
}
