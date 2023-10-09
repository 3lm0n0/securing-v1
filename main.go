package main

import (
	"net/http"
	""github.com/3lm0n0/securing-v1/auth""
)

func main() {
	http.Handle("/register", http.HandlerFunc(auth.RegisterHandler))
	http.Handle("/login", http.HandlerFunc(auth.LoginHandler))
	http.Handle("/secureEndpoint", auth.AuthMiddleware(http.HandlerFunc(secureEndpointHandler)))

	http.ListenAndServe(":8080", nil)
}

func secureEndpointHandler(w http.ResponseWriter, r *http.Request) {
	// Implement secure endpoint logic here
}