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

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodDelete:
		{
			fmt.Println("I am here")
			deleteUser(w, r)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
	// Retrieve form values from the request
	username := r.FormValue("username")
	email := r.FormValue("email")

	// Create a new User object with the form values
	newUser := User{
		ID:       len(users) + 1,
		Username: username,
		Email:    email,
	}

	// Log the newUser object
	fmt.Println("New user:", newUser)

	// Add the newUser to the users slice
	users = append(users, newUser)

	// Respond with the newUser as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extracting the "id" parameter from the request URL
	id := r.URL.Query().Get("id")
	fmt.Println("ID:", id)

	// Assuming a simple case where IDs are integers
	// In a real-world scenario, you would need proper validation
	// and error handling here
	userID := 0
	fmt.Sscanf(id, "%d", &userID)

	// Find the index of the user with the specified ID
	index := -1
	for i, user := range users {
		if user.ID == userID {
			index = i
			break
		}
	}

	if index == -1 {
		http.NotFound(w, r)
		return
	}

	// Remove the user from the slice
	users = append(users[:index], users[index+1:]...)

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User deleted successfully"}
	json.NewEncoder(w).Encode(response)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file containing the basic UI
	http.ServeFile(w, r, "./index.html")
}

func main() {
	// Initializing some dummy users for demonstration purposes
	users = append(users, User{ID: 1, Username: "user1", Email: "user1@example.com"})
	users = append(users, User{ID: 2, Username: "user2", Email: "user2@example.com"})

	http.HandleFunc("/", mainPage)

	// API endpoints
	http.HandleFunc("/api/users", getUsers)
	http.HandleFunc("/api/user", handleUser)

	// Start the server
	fmt.Println("API server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
