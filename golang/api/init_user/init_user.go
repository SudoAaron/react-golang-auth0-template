package init_user

import (
	"database/sql"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func Handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		// Check if user exists in the database
		exists, err := UserExists(db, auth0UserID)
		if err != nil {
			// Handle database error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Internal Server Error"}`))
			return
		}

		if !exists {
			// User does not exist, set up user in DB with default settings
			err := SetupUser(db, auth0UserID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Failed to set up user"}`))
				return
			}
		}

		if exists {
			// User does not exist, set up user in DB with default settings
			err := UpdateLastLogin(db, auth0UserID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Failed to update last_login"}`))
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UserExists checks if a user with the given auth0UserID exists in the database
func UserExists(db *sql.DB, auth0UserID string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE auth0_user_id = $1)"
	var exists bool
	err := db.QueryRow(query, auth0UserID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// SetupUser sets up a new user in the database with default settings
func SetupUser(db *sql.DB, auth0UserID string) error {
	query := `
        INSERT INTO users (auth0_user_id, last_login, registration_date)
        VALUES ($1, NOW(), NOW())
    `
	_, err := db.Exec(query, auth0UserID)
	if err != nil {
		return err
	}

	// Assign the "user" role to the new user
	roleQuery := `
        INSERT INTO user_roles (user_id, role_id)
        SELECT u.id, r.id
        FROM users u
        CROSS JOIN roles r
        WHERE u.auth0_user_id = $1
        AND r.role_name = 'user'
    `
	_, err = db.Exec(roleQuery, auth0UserID)
	if err != nil {
		return err
	}

	return err
}

// UpdateLastLogin updates the last_login time for the user
func UpdateLastLogin(db *sql.DB, auth0UserID string) error {
	query := `
        UPDATE users
        SET last_login = NOW()
        WHERE auth0_user_id = $1
    `
	_, err := db.Exec(query, auth0UserID)
	return err
}
