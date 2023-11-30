package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = CreateTablesIfNotExists(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Wrapper function for all DB setup
func CreateTablesIfNotExists(db *sql.DB) error {
	err := CreateUsersTableIfNotExists(db)
	if err != nil {
		return err
	}

	err = CreateRolesTableIfNotExists(db)
	if err != nil {
		return err
	}

	err = CreateUserRolesTableIfNotExists(db)
	if err != nil {
		return err
	}

	err = AutoPopulateRoles(db)
	if err != nil {
		return err
	}

	return nil
}

// Set up users Table if it doesn't exist.
func CreateUsersTableIfNotExists(db *sql.DB) error {
	createUsersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            auth0_user_id TEXT NOT NULL,
            last_login TIMESTAMP,
            registration_date TIMESTAMP NOT NULL
        );
    `
	_, err := db.Exec(createUsersTable)
	if err != nil {
		return err
	}
	return nil
}

// Set up roles Table if it doesn't exist.
func CreateRolesTableIfNotExists(db *sql.DB) error {
	createRolesTable := `
        CREATE TABLE IF NOT EXISTS roles (
            id SERIAL PRIMARY KEY,
            role_name VARCHAR(255) NOT NULL UNIQUE
        );
    `
	_, err := db.Exec(createRolesTable)
	if err != nil {
		return err
	}
	return nil
}

// AutoPopulateRoles populates the "roles" table with default roles if it's empty.
func AutoPopulateRoles(db *sql.DB) error {
	// Check if the "roles" table is empty
	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM roles").Scan(&rowCount)
	if err != nil {
		return err
	}

	// If the table is empty, insert default roles
	if rowCount == 0 {
		defaultRoles := []string{"user", "admin"}
		for _, roleName := range defaultRoles {
			_, err := db.Exec("INSERT INTO roles (role_name) VALUES ($1)", roleName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Set up user_roles Table if it doesn't exist.
func CreateUserRolesTableIfNotExists(db *sql.DB) error {
	createUserRolesTable := `
        CREATE TABLE IF NOT EXISTS user_roles (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL,
            role_id INT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users(id),
            FOREIGN KEY (role_id) REFERENCES roles(id)
        );
    `
	_, err := db.Exec(createUserRolesTable)
	if err != nil {
		return err
	}
	return nil
}
