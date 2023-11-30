package middleware

import (
	"net/http"
	"strings"

	"database/sql"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/lib/pq"
)

func CheckRoles(allowedRoles []string, db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// func CheckRoles(allowedRoles []string, db *sql.DB, next http.Handler) http.Handler {
			// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Get the token claims from the request context
			token := r.Context().Value(jwtmiddleware.ContextKey{})
			if token == nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"message":"Unauthorized."}`))
				return
			}

			claims, ok := token.(*validator.ValidatedClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"message":"Unauthorized."}`))
				return
			}

			sub := claims.RegisteredClaims.Subject

			subParts := strings.Split(sub, "|")
			if len(subParts) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"message":"Unauthorized."}`))
				return
			}

			auth0UserID := subParts[1]

			// Check if the user has any of the allowed roles
			if !UserHasAllowedRole(db, auth0UserID, allowedRoles) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message":"Access Denied."}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// UserHasAllowedRole checks if the user has any of the allowed roles
func UserHasAllowedRole(db *sql.DB, auth0UserID string, allowedRoles []string) bool {
	// Query the database to check if the user has any of the allowed roles
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			JOIN roles r ON ur.role_id = r.id
			JOIN users u ON ur.user_id = u.id
			WHERE u.auth0_user_id = $1
			AND r.role_name = ANY($2)
		)
	`

	var exists bool
	err := db.QueryRow(query, auth0UserID, pq.Array(allowedRoles)).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
