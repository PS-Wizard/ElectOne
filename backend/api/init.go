package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var RDB *redis.Client
var DB *sql.DB
var CTX = context.Background()

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

func InitializeRedis(host, port, password string) {
	addr := fmt.Sprintf("%s:%s", host, port)
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := RDB.Ping(CTX).Result()
	if err != nil {
		log.Fatalf("Failed To Connect To Redis: %v", err)
	}
	log.Println("Redis Connection Established and publisher started")
}

func CloseRedis() {
	if RDB != nil {
		RDB.Close()
		log.Println("Redis Connection Closed")
	}
}
