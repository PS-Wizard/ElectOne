package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID        int    `json:"user_id"`
	CitizenshipID string `json:"citizenship_id"`
	VoterCardID   string `json:"voter_card_id"`
	Password      string `json:"password"`
	TotpSecret    string `json:"totp_secret"`
	Photos        string `json:"photos"`
}

func CreateUser(user *User) (int, error) {
	query := `
		INSERT INTO User (CitizenshipID, VoterCardID, Password, TotpSecret, Photos)
		VALUES (?, ?, ?, ?, ?)`

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("Failed To Bcrypt User Password: %v", err)
	}
	_, err = db.DB.Exec(query, user.CitizenshipID, user.VoterCardID, hashedPass, user.TotpSecret, user.Photos)
	if err != nil {
		return 0, err
	}

	var userID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func UpdateUser(userID int, updatedUser *User) error {
	query := `
		UPDATE User
		SET CitizenshipID = ?, VoterCardID = ?, Password = ?, TotpSecret = ?, Photos = ?
		WHERE UserID = ?`

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Failed To Bcrypt User Password: %v", err)
	}
	_, err = db.DB.Exec(query, updatedUser.CitizenshipID, updatedUser.VoterCardID, hashedPass, updatedUser.TotpSecret, updatedUser.Photos, userID)
	return err
}

func GetUsers(limit, offset int) ([]User, error) {
	query := `
		SELECT UserID, CitizenshipID, VoterCardID, Password, TotpSecret, Photos
		FROM User
		ORDER BY UserID DESC
		LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.CitizenshipID, &user.VoterCardID, &user.Password, &user.TotpSecret, &user.Photos); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(userID int) (*User, error) {
	query := `
		SELECT UserID, CitizenshipID, VoterCardID, Password, TotpSecret, Photos
		FROM User
		WHERE UserID = ?`
	var user User
	err := db.DB.QueryRow(query, userID).Scan(&user.UserID, &user.CitizenshipID, &user.VoterCardID, &user.Password, &user.TotpSecret, &user.Photos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No User With That ID")
		}
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userID int) error {
	query := `DELETE FROM User WHERE UserID = ?`
	_, err := db.DB.Exec(query, userID)
	return err
}

func GetUserByCitizenAndVoterID(citizenID, voterID string) (*User, error) {
	query := `SELECT * FROM User WHERE CitizenshipID = ? AND VoterCardID = ?`
	row := db.DB.QueryRow(query, citizenID, voterID)

	var user User
	err := row.Scan(&user.UserID, &user.CitizenshipID, &user.VoterCardID, &user.Password, &user.TotpSecret, &user.Photos)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SetUserTOTPSecret(userID int, secret string) error {
	query := `UPDATE User SET TotpSecret = ? WHERE UserID = ?`
	_, err := db.DB.Exec(query, secret, userID)
	return err
}

func UpdateUserPassword(userID int, newPassword string) error {
	query := `
		UPDATE User
		SET Password = ?
		WHERE UserID = ?`

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Failed To Bcrypt User Password: %v", err)
	}
	_, err = db.DB.Exec(query, hashedPass, userID)
	return err
}
