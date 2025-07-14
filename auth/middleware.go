package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type contextKey string

const userKey = contextKey("user")

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Missing token cookie", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		username := claims["username"].(string)

		ctx := context.WithValue(r.Context(), userKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
