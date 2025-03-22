package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	authToken, dbURl := os.Getenv("TURSO_AUTH_TOKEN"), os.Getenv("TURSO_DATABASE_URL")
	if authToken == "" || dbURl == "" {
		log.Fatal("TURSO_AUTH_TOKEN / TURSO_DATABASE_URL environment variable is required")
	}
	url := fmt.Sprintf("%s?authToken=%s", dbURl, authToken)
	fmt.Printf("Trying To Establish A Connection To: %s\n", dbURl)
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s\n", url, err)
		os.Exit(1)
	}
	fmt.Println("Successfully Established A Connection")
	DB = db
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database Connection Closed")
	}
}
