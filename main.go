package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net"
	"net/http"
	"strings"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // replace with your Redis server address
		Password: "",               // replace with your password if you have one
		DB:       0,                // use default DB
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		w.Write([]byte("Hello, your IP is: " + ip))

		err := rdb.Set(ctx, ip, ip, 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, ip).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(ip, val)
	})

	http.ListenAndServe(":8080", nil)
}
