package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net"
	"net/http"
	"strings"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // replace with your Redis server address
		Password: "",               // replace with your password if you have one
		DB:       0,                // use default DB
	})

	//rdb.Set(ctx, "key", "value", 0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		w.Write([]byte("Hello, your IP is: " + ip))

		err := rdb.Set(ctx, ip, "1", time.Second*10).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, ip).Result()
		if err != nil {
			panic(err)
		}

		rdb.Incr(ctx, ip)
		rdb.Incr(ctx, ip)
		rdb.Incr(ctx, ip)
		rdb.Incr(ctx, ip)
		fmt.Println(ip, val)
	})

	http.ListenAndServe(":8080", nil)
}
