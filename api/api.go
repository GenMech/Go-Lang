package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents a basic user structure
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extracting the "id" parameter from the request URL
	id := r.URL.Query().Get("id")
	fmt.Println("ID:", id)

	// Assuming a simple case where IDs are integers
	// In a real-world scenario, you would need proper validation
	// and error handling here
	userID := 0
	fmt.Sscanf(id, "%d", &userID)

	// Find the user with the specified ID
	var foundUser *User
	for _, user := range users {
		if user.ID == userID {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(foundUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Generate a simple ID for the new user
	newUser.ID = len(users) + 1

	users = append(users, newUser)

	json.NewEncoder(w).Encode(newUser)
}

func main() {
	// Initializing some dummy users for demonstration purposes
	users = append(users, User{ID: 1, Username: "user1", Email: "user1@example.com"})
	users = append(users, User{ID: 2, Username: "user2", Email: "user2@example.com"})

	// API endpoints
	http.HandleFunc("/api/users", getUsers)
	http.HandleFunc("/api/user", getUser)
	http.HandleFunc("/api/user", createUser)

	// Start the server
	fmt.Println("API server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
