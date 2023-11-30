package get_roles

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func Handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		userID, err := GetUserID(db, auth0UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Internal Server Error"}`))
			return
		}
		if userID != 0 {
			roles, err := GetUserRoles(db, userID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Failed to retrieve user roles"}`))
				return
			}

			response := struct {
				Roles []string `json:"roles"`
			}{
				Roles: roles,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Failed to encode response"}`))
				return
			}
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
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

func GetUserRoles(db *sql.DB, userID int) ([]string, error) {
	query := `
        SELECT r.role_name
        FROM user_roles ur
        JOIN roles r ON ur.role_id = r.id
        WHERE ur.user_id = $1
    `
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string

	for rows.Next() {
		var roleName string
		err := rows.Scan(&roleName)
		if err != nil {
			return nil, err
		}
		roles = append(roles, roleName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
