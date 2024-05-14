package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisDb struct {
	client *redis.Client
}

func (rd *RedisDb) Set(ctx context.Context, key, value string, timeout time.Duration) error {
	return rd.client.Set(ctx, key, value, timeout).Err()
}

func (rd *RedisDb) Get(ctx context.Context, key string) (string, error) {
	return rd.client.Get(ctx, key).Result()
}

func (rd *RedisDb) Incr(ctx context.Context, key string) error {
	return rd.client.Incr(ctx, key).Err()
}

func NewRedisDb(addr, password string, db int) (*RedisDb, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisDb{client: client}, nil
}
