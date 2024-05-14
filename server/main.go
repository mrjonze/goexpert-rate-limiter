package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mrjonze/goexpert/rate-limiter/pkg"
	"net/http"
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
	return pkg.RateLimiteHandler(next)
}
