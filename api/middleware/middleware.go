package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	sugar "sugar/data"
	"sugar/globals/auth"
	"sugar/helpers/response"

	sqlite3 "github.com/mattn/go-sqlite3"
)

type Middleware struct {
	queries *sugar.Queries
}

func NewMiddleware(queries *sugar.Queries) *Middleware {
	return &Middleware{queries: queries}
}

func getAuthorizationHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing_header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid_header")
	}

	return parts[1], nil
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sessionId, err := getAuthorizationHeader(r)
		if err != nil {
			switch err.Error() {
			case "missing_header":
				response.BadRequest(w, "Authorization header is missing.")
				return
			case "invalid_header":
				response.BadRequest(w, "Authorization header is invalid.")
				return
			default:
				slog.Debug("Unknown authorization error", slog.Any("error", err))
			}

			response.InternalServerError(w, err, "Something went wrong")
			return
		}

		session, err := m.queries.GetSessionByID(r.Context(), &sessionId)
		if err != nil {
			var sqliteErr sqlite3.Error
			if errors.As(err, &sqliteErr) {
				if sqliteErr.Code == sqlite3.ErrNotFound {
					response.Conflict(w, "Session not found")
					return
				}
			}

			response.InternalServerError(w, err, "Something went wrong")
			return
		}

		ctx := context.WithValue(r.Context(), auth.SessionKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
