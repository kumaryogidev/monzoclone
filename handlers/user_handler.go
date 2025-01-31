package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"monzo-like-bank/models"
	"monzo-like-bank/utils"

	"github.com/gocql/gocql"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Generate a UUID for the new user
	user.ID = gocql.TimeUUID()
	user.CreatedAt = time.Now()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Password hashing failed")
		return
	}
	user.Password = hashedPassword

	// Insert user into Cassandra
	err = utils.Session.Query(`INSERT INTO users (id, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)`,
		user.ID, user.Name, user.Email, user.Password, user.CreatedAt).Exec()

	if err != nil {
		log.Printf("Failed to insert user: %v\n", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// Query to fetch all users
	iter := utils.Session.Query("SELECT id, name, email, created_at FROM users").Iter()
	var user models.User

	for iter.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		log.Println("Error fetching users:", err)
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	// Return the list of users
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
