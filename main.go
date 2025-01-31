package main

import (
	"fmt"
	"log"
	"monzo-like-bank/handlers"
	"monzo-like-bank/utils"
	"net/http"
)

func main() {
	utils.ConnectCassandra() // Connect to Cassandra DB
	http.HandleFunc("/api/health", HealthCheckHandler)
	http.HandleFunc("/api/register", handlers.RegisterUser)
	http.HandleFunc("/api/users", handlers.ListUsers)

	port := "8080"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"OK"}`))
}
