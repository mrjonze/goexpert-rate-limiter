package main

import (
	"fmt"
	"net/http"
)

func main() {
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
