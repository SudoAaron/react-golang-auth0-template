package set_role

import (
	"database/sql"
	"fmt"
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

		// User does not exist, set up user in DB with default settings
		err := MakeUserAdmin(db, auth0UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Failed to update set_role"}`))
			return
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

// MakeUserAdmin adds the "admin" role to the user
func MakeUserAdmin(db *sql.DB, auth0UserID string) error {
	// Check if the user exists in the users table

	userID, err := GetUserID(db, auth0UserID)
	if err != nil {
		return err
	}

	if userID == 0 {
		return fmt.Errorf("user not found")
	}

	// Check if the "admin" role exists in the roles table
	roleID, err := GetRoleIDByName(db, "admin")
	if err != nil {
		return err
	}

	if roleID == 0 {
		return fmt.Errorf("admin role not found")
	}

	// Check if the user already has the "admin" role
	roleExists, err := UserHasRole(db, userID, roleID)
	if err != nil {
		return err
	}

	if !roleExists {
		// User does not have the "admin" role, insert it into the user_roles table
		insertQuery := "INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)"
		_, err := db.Exec(insertQuery, userID, roleID)
		if err != nil {
			return err
		}
	}

	return nil
}

// UserHasRole checks if the user has a specific role
func UserHasRole(db *sql.DB, userID int, roleID int) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM user_roles WHERE user_id = $1 AND role_id = $2)"
	var roleExists bool
	err := db.QueryRow(query, userID, roleID).Scan(&roleExists)
	if err != nil {
		return false, err
	}
	return roleExists, nil
}

func GetUserID(db *sql.DB, auth0UserID string) (int, error) {
	query := "SELECT id FROM users WHERE auth0_user_id = $1"
	var id int
	err := db.QueryRow(query, auth0UserID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetRoleIDByName retrieves the role_id based on the role_name
func GetRoleIDByName(db *sql.DB, roleName string) (int, error) {
	query := "SELECT id FROM roles WHERE role_name = $1"
	var roleID int
	err := db.QueryRow(query, roleName).Scan(&roleID)
	if err != nil {
		return 0, err
	}
	return roleID, nil
}
