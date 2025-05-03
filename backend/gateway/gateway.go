package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

var servers = []string{
	"http://localhost:5173",
	// "http://localhost:3002",
	// "http://localhost:3003",
}

var counter uint64

// nextServer determines the next server in a round-robin fashion.
func nextServer() string {
	i := atomic.AddUint64(&counter, 1)
	return servers[i%uint64(len(servers))]
}

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	// Get the next server to redirect to
	// target := nextServer()

	// Serve a basic HTML page with a loading animation
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Loading...</title>
			<style>
				body {
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
					background-color: #f4f4f9;
					margin: 0;
					flex-direction: column;
				}
				.spinner {
					width: 50px;
					height: 50px;
					border: 6px solid rgba(255, 255, 255, 0.3);
					border-top: 6px solid #3498db;
					border-radius: 50%;
					animation: spin 1s linear infinite;
				}
				@keyframes spin {
					0% { transform: rotate(0deg); }
					100% { transform: rotate(360deg); }
				}
				.text {
					font-size: 1.2rem;
					color: #555;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="spinner"></div>
			<p class="text">Finding the best server for you...</p>
		</body>
		<script>
		setTimeout(()=>{
			window.location.href = "http://localhost:5173/user/login"
		},2000);
		</script>
		</html>
	`

	// Simulate "doing work" by sleeping for a bit before the redirect
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))

	// http.Redirect(w, r, target, http.StatusFound)
}

func main() {
	http.HandleFunc("/", gatewayHandler)

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
