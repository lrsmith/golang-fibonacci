package middleware

/*
Taken from https://github.com/gorilla/mux#middleware

Ideally authentication information would not be hard-coded in code.
Instead the middle-ware function would call out to a service such as vault, etc.
*/

import (
	"log"
	"net/http"
)

type AuthenticationMiddleware struct {
	TokenUsers map[string]string
}

// Initialize it somewhere
func (amw *AuthenticationMiddleware) Populate() {

	amw.TokenUsers = make(map[string]string)

	amw.TokenUsers["00000000"] = "zaphod"
	amw.TokenUsers["aaaaaaaa"] = "ford"
}

// Middleware function, which will be called for each request
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.TokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", 403)
		}
	})
}
