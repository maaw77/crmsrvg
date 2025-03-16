package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maaw77/crmsrvg/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenSting := strings.Split(r.Header.Get("Authorization"), " ")
		if len(tokenSting) < 2 || tokenSting[1] == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(`{"deatails": "token is not provided"}`)
			return
		}

		token, err := auth.VerifyToken(strings.TrimSpace(tokenSting[1]))
		if token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			json.NewEncoder(w).Encode(`{"deatails": "that's not even a token"}`)
		case errors.Is(err, jwt.ErrTokenExpired):
			json.NewEncoder(w).Encode(`{"deatails": "token is e expired"}`)
		default:
			json.NewEncoder(w).Encode(`{"deatails": "couldn't handle this token"}`)
		}

	})
}
