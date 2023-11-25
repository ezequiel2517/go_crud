package middlewares

import (
	"api/api/authentication"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JWTKey = []byte(os.Getenv("MY_VARIABLE"))

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Split(authHeader, " ")[1]
		token, err := jwt.ParseWithClaims(tokenStr, &authentication.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JWTKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
