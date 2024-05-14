package pkg

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/mrjonze/goexpert/rate-limiter/db"
	"github.com/mrjonze/goexpert/rate-limiter/server/config"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RateLimiteHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		configs, err := config.LoadConfig()
		if err != nil {
			panic(err)
		}

		tokenName := r.Header.Get("API_KEY")
		ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

		var key string
		var expiration int
		var limit int

		if tokenName != "" && tokenName != configs.TokenName {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
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
			panic(err)
		}

		isBlocked, _ := db.Get(ctx, "b-"+key)

		if isBlocked != "" {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
			return
		}

		hitsStr, err := db.Get(ctx, key)

		if errors.Is(err, redis.Nil) {
			db.Set(ctx, key, "1", time.Duration(expiration)*time.Second)
		} else {
			hits, err := strconv.Atoi(hitsStr)
			if err != nil {
				panic(err)
			}
			if hits >= limit {
				db.Set(ctx, "b-"+key, "1", time.Duration(expiration)*time.Second)
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
				return
			} else {
				db.Incr(ctx, key)
			}
		}

		next.ServeHTTP(w, r)
	})
}
