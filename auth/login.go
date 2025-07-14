package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	println("HI")
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil || creds.Username != "admin" || creds.Password != "whoami" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	w.Write([]byte("Logged in with cookie"))
}
