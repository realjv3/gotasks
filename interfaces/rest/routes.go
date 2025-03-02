package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func (h *UserHandler) RegisterUserRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.createUser)
			r.Get("/{id}", h.getUser)
		})
	})
}

func (h *AuthHandler) RegisterAuthRoutes(r chi.Router) {
	r.Post("/login", h.Login)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		token, err := jwt.Parse(header[7:], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_KEY")), nil
		})

		switch {
		case token.Valid:
		case errors.Is(err, jwt.ErrTokenMalformed):
			http.Error(w, "Malformed token", http.StatusUnauthorized)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}
