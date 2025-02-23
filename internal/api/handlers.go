package api

import (
	"net/http"
	"log"
)
// Middleware to check if the user is authenticated
func isAuthenticated(r *http.Request) bool {
	// session, r := store.Get(r, "session")
	// authenticated, ok := session.Values["authenticated"].(bool)
	// return ok && authenticated
	return true
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Login request")	
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
  // Validate the request if needed
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
  // Process the request 
	 

  // Return the response
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func DescribeTaskHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}
