package main

import (
	"fmt"
	"net/http"

	"github.com/PS-Wizard/VotingSystem-API/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.HandleAPIRoutes(mux)
	fmt.Println("Listening On Port :8080")
	http.ListenAndServe(":8080", mux)

}
