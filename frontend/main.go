package main

import (
	"fmt"
	"net/http"

	"github.com/PS-Wizard/Plugs/controllers/routes"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	routes.HandleRoutes(mux)
	fmt.Println("Server Listening at :8080")
	http.ListenAndServe(":8080", mux)
}
