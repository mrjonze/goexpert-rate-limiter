package main

import (
	"fmt"
	"github.com/mrjonze/goexpert/rate-limiter/server/config"
	"net/http"
)

func main() {
	configs, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	fmt.Println("Token name: ", configs.TokenName)

	http.HandleFunc("/", readHeader)
	http.ListenAndServe(":8080", nil)
}

func readHeader(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}
