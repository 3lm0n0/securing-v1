package auth

import (
	"fmt"
	"net/http"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Handler")
	// Implement user registration logic here (store user in database)
	// ...
	// After successful registration, generate and send JWT token to client
	
	tokenString := generateJWTToken(user)

	w.Write([]byte(tokenString))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Handler")
	// Implement user login logic here (validate user credentials and generate JWT token)
	// ...
	// After successful login, generate and send JWT token to client
	
	tokenString := generateJWTToken(user)

	w.Write([]byte(tokenString))
}