package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Middleware to check if the user is authenticated
func isAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	// Get the claims after the token is validated
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return false
	}
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}

	claims, err := ParseJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return false
	}

	// If token is valid, you can access the claims (for example, print the username)
	w.Write([]byte(fmt.Sprintf("Hello %s, you are authorized!", claims.Username)))
	return true
}

// Struct for holding the login credentials
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	username := req.Username
	// password := req.Password

	log.Printf("Login request: %s", username)

	// Generate the JWT token
	tokenString, err := GenerateJWT(username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send the token to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, tokenString) + "\n"))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request if needed
	if !isAuthenticated(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// Process the request

	// Return the response
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func DescribeTaskHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}
