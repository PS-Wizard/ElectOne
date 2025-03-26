package users

import (
	"database/sql"
	"fmt"
	"github.com/PS-Wizard/ElectOneAPI/api"
)

// Fetch a user by ID from the database
func getUser(userID string) (*user, error) {
	var u user
	query := "SELECT userID, citizenID, password FROM users WHERE userID = ?"
	row := api.DB.QueryRow(query, userID)
	err := row.Scan(&u.UserID, &u.CitizenID, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	return &u, nil
}

// Get users with pagination
func getUsersPaginated(offset int) ([]user, error) {
	var users []user
	query := "SELECT userID, citizenID, password FROM users LIMIT 10 OFFSET ?"
	rows, err := api.DB.Query(query, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.UserID, &u.CitizenID, &u.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}
	return users, nil
}

// Create a new user in the database
func createNewUser(u user) error {
	query := "INSERT INTO users (citizenID, password) VALUES (?, ?)"
	_, err := api.DB.Exec(query, u.CitizenID, u.Password)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

// Update user details in the database
func updateUserDetails(u user, userID string) error {
	query := "UPDATE users SET citizenID = ?, password = ? WHERE userID = ?"
	_, err := api.DB.Exec(query, u.CitizenID, u.Password, userID)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

// Delete a user from the database
func deleteUser(userID string) error {
	query := "DELETE FROM users WHERE userID = ?"
	_, err := api.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
