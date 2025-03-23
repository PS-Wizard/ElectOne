package main

import (
	"fmt"
	"net/http"

	"github.com/PS-Wizard/VotingSystem/db"
	"github.com/PS-Wizard/VotingSystem/routes"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	mux := http.NewServeMux()
	routes.HandleRoutes(mux)
	fmt.Println("Server Listening On :8080")
	http.ListenAndServe(":8080", mux)
}
