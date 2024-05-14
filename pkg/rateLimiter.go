package pkg

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/mrjonze/goexpert/rate-limiter/db"
	"github.com/mrjonze/goexpert/rate-limiter/server/config"
	"net/http"
	"strconv"
	"time"
)

func RateLimitRequests(ctx context.Context, ip string, tokenName string) (string, int) {
	configs, err := config.LoadConfig()
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	var key string
	var expiration int
	var limit int

	if tokenName != "" && tokenName != configs.TokenName {
		return "invalid token", http.StatusUnauthorized
	}
	if tokenName == configs.TokenName {
		key = tokenName
		expiration = configs.BlockTimeToken
		limit = configs.RequestLimitToken
	} else {
		key = ip
		expiration = configs.BlockTimeIp
		limit = configs.RequestLimitIp
	}

	db, err := db.NewRedisDb(configs.DatabaseUrl, "", 0)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	isBlocked, _ := db.Get(ctx, "b-"+key)

	if isBlocked != "" {
		return "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests
	}

	hitsStr, err := db.Get(ctx, key)

	if errors.Is(err, redis.Nil) {
		db.Set(ctx, key, "1", time.Duration(expiration)*time.Second)
	} else {
		hits, err := strconv.Atoi(hitsStr)
		if err != nil {
			return err.Error(), http.StatusInternalServerError
		}
		if hits >= limit {
			db.Set(ctx, "b-"+key, "1", time.Duration(expiration)*time.Second)
			return "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests
		} else {
			db.Incr(ctx, key)
		}
	}
	return "ok", http.StatusOK
}
