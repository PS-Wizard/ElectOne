package users

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/ElectOneAPI/api"
	"golang.org/x/crypto/bcrypt"
)

// Fetch a user by ID from the database
func getUser(userID string) (*User, error) {
	var u User
	query := "SELECT userID, citizenID, password, phonenumber, tag FROM users WHERE userID = ?"
	row := api.DB.QueryRow(query, userID)
	err := row.Scan(&u.UserID, &u.CitizenID, &u.Password, &u.Phone, &u.Tag)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	return &u, nil
}

// Get users with pagination
func getUsersPaginated(offset int) ([]User, error) {
	var users []User
	query := "SELECT userID, citizenID, password, phonenumber, tag FROM users LIMIT 10 OFFSET ?"
	rows, err := api.DB.Query(query, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.UserID, &u.CitizenID, &u.Password, &u.Phone, &u.Tag)
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
func CreateNewUser(u User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	query := "INSERT INTO users (citizenID, password, phonenumber, tag) VALUES (?, ?, ?, ?)"
	_, err = api.DB.Exec(query, u.CitizenID, string(hashedPassword), u.Phone, u.Tag)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

// Update user details in the database
func updateUserDetails(u User, userID string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	query := "UPDATE users SET citizenID = ?, password = ?, phonenumber = ?, tag = ? WHERE userID = ?"
	_, err = api.DB.Exec(query, u.CitizenID, string(hashedPassword), u.Phone, u.Tag, userID)
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
