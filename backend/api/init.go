package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitializeDB(authToken, dbUrl string) {
	url := fmt.Sprintf("%s?authToken=%s", dbUrl, authToken)
	log.Printf("Trying To Establish A Connection To %s", dbUrl)
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed To Establish A Connection To %s: %v", dbUrl, err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	log.Println("Connection Established To TURSO")

	DB = db
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Fatalf("Connection To Database Closed.")
	}
}
