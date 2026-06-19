package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("-> %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("<- %s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "test" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return // chain stops here, next is never called
		}
		next.ServeHTTP(w, r)
	})
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, "service unavailable", http.StatusServiceUnavailable)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/public", chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "public endpoint — no auth required")
		}),
		recovery, logging,
	))

	mux.Handle("/private", chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "private endpoint — auth passed")
		}),
		recovery, logging, auth,
	))

	mux.Handle("/panic", chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("something went wrong")
		}),
		recovery, logging,
	))

	log.Println("listening on :9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
