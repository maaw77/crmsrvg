package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maaw77/crmsrvg/internal/auth"
)

// ErrorMessage is a detailed error message.
type errorMessage struct {
	// Description of the situation
	// example: An error occurred
	Details string `json:"details"`
}

// AuthMiddleware is an authentication middleware that verifies the session token..
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthMiddleware Serving:", r.URL.Path, "from", r.Host)
		tokenSting := strings.Split(r.Header.Get("Authorization"), " ")
		if len(tokenSting) < 2 || tokenSting[1] == "" {
			log.Println("token is not provided")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorMessage{Details: "token is not provided"})
			return
		}

		// log.Printf("token = %s\n", tokenSting[1])
		token, err := auth.VerifyToken(strings.TrimSpace(tokenSting[1]))

		if token != nil && token.Valid {
			log.Println("token is valid")
			next.ServeHTTP(w, r)
			return

		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			log.Println("that's not even a token")
			json.NewEncoder(w).Encode(errorMessage{Details: "that's not even a token"})
		case errors.Is(err, jwt.ErrTokenExpired):
			log.Println("token is expired")
			json.NewEncoder(w).Encode(errorMessage{Details: "token is expired"})
		default:
			log.Println("couldn't handle this token")
			json.NewEncoder(w).Encode(errorMessage{Details: "couldn't handle this token"})

		}
	})
}
