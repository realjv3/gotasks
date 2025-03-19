package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func (h *UserHandler) RegisterUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.createUser)

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware)
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
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("JWT_KEY"), nil
		})

		switch {
		case token.Valid:
		case errors.Is(err, jwt.ErrTokenMalformed):
			http.Error(w, "malformed token", http.StatusUnauthorized)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			http.Error(w, "invalid signature", http.StatusUnauthorized)
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		sub, err := token.Claims.GetSubject()
		if err != nil {
			http.Error(w, "failed to get subject from JWT", http.StatusInternalServerError)
			return
		}

		userID, err := strconv.Atoi(sub)
		if err != nil {
			http.Error(w, "error converting user id", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "user-id", userID)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *TaskHandler) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", h.createTask)
			r.Get("/{userID}", h.getTasksByUser)
			r.Put("/finish/{taskID}", h.finishTask)
		})
	})
}
