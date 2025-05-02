package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	AdminID    int    `json:"admin_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	TotpSecret string `json:"totp_secret"`
}

func CreateAdmin(admin *Admin) (int, error) {
	query := `
		INSERT INTO Admin (Email, Password, TotpSecret)
		VALUES (?, ?, ?)`

	hashed, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("Error Bcrypting Password: %v", err)
	}
	_, err = db.DB.Exec(query, admin.Email, hashed, admin.TotpSecret)
	if err != nil {
		return 0, err
	}

	var adminID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&adminID)
	if err != nil {
		return 0, err
	}

	return adminID, nil
}

func UpdateAdmin(adminID int, updatedAdmin *Admin) error {
	query := `
		UPDATE Admin
		SET Email = ?, Password = ?, TotpSecret = ?
		WHERE AdminID = ?`
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(updatedAdmin.Password), bcrypt.DefaultCost)
	_, err = db.DB.Exec(query, updatedAdmin.Email, hashedPwd, updatedAdmin.TotpSecret, adminID)
	return err
}

func GetAdmins(limit, offset int) ([]Admin, error) {
	query := `
		SELECT AdminID, Email, Password, TotpSecret
		FROM Admin
		ORDER BY AdminID DESC
		LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []Admin
	for rows.Next() {
		var admin Admin
		if err := rows.Scan(&admin.AdminID, &admin.Email, &admin.Password, &admin.TotpSecret); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}

func GetAdminByID(adminID int) (*Admin, error) {
	query := `
		SELECT AdminID, Email, Password, TotpSecret
		FROM Admin
		WHERE AdminID = ?`
	var admin Admin
	err := db.DB.QueryRow(query, adminID).Scan(&admin.AdminID, &admin.Email, &admin.Password, &admin.TotpSecret)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Admin With That ID")
		}
		return nil, err
	}
	return &admin, nil
}

func GetAdminByEmail(email string) (*Admin, error) {
	query := `SELECT AdminID, Email, Password, TotpSecret FROM Admin WHERE Email = ?`
	row := db.DB.QueryRow(query, email)

	var admin Admin
	err := row.Scan(&admin.AdminID, &admin.Email, &admin.Password, &admin.TotpSecret)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}
