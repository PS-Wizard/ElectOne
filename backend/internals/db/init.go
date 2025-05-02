package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

const url string = "libsql://electone-wizard.aws-ap-south-1.turso.io"
const token string = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3NDYxNzcyMjAsImlkIjoiOTJlNDE5NDAtZDkyZS00MzAwLTk2ZmYtNDZhZjQxODkzOGYyIiwicmlkIjoiMTUxZjI3MWUtY2U2ZC00MDQwLWE3ZTYtYjJmNzI2OTJmYWVhIn0.4PizC8s2d1-8pWw1n8K21_HCDK0hcTgJpMsQNCzy_aFKD9CTOUHfZNzJByEKRzuciDC5CYmf-cJ6oTrkI9q6Cw"

func InitTurso() {
	db, err := sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", url, token))
	if err != nil {
		log.Fatalf("Error Connecting To Turso %v", err)
	}
	DB = db
}

func CloseTurso() {
	if DB != nil {
		DB.Close()
	}
}
