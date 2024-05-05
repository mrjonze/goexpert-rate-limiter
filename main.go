package main

import (
	"net"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		w.Write([]byte("Hello, your IP is: " + ip))
	})

	http.ListenAndServe(":8080", nil)
}
