package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware")
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Check if user has required permissions (based on roles, etc.)
		// ...

		next.ServeHTTP(w, r)
	})
}

func generateJWTToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodH256, jwt.MapClaims {
		"id": user.ID,
		"email": user.Email,
		"role": user.Role,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error generating token: ", err)
		return ""
	}

	return tokenString
}