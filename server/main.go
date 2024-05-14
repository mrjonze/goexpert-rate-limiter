package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mrjonze/goexpert/rate-limiter/pkg"
	"net"
	"net/http"
	"strings"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(rateLimiter)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", r)
}

func rateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenName := r.Header.Get("API_KEY")
		ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

		message, status := pkg.RateLimitRequests(ctx, ip, tokenName)
		w.WriteHeader(status)
		w.Write([]byte(message))

		next.ServeHTTP(w, r)
	})
}
