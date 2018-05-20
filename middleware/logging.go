package middleware

// Taken from https://github.com/gorilla/mux#middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware ...
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
