package main

import (
	"log"
	"net/http/cookiejar"
	"net/http"
	"fmt"
	"bytes"
)

func main() {
	log.Println("client running")

	// Set up an HTTP client with cookie support
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Failed to create cookie jar: %v", err)
	}
	client := &http.Client{
		Jar: jar, // Enable the cookie jar for storing cookies
	}

	// Step 1: Login
	username := "admin"
	// password := "password"
	loginData := fmt.Sprintf("username=%s", username)
	loginRequest, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", bytes.NewBufferString(loginData))
	if err != nil {
		log.Fatalf("Failed to create login request: %v", err)
	}
	loginRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Step 2: Send login request
	response, err := client.Do(loginRequest)
	if err != nil {
		log.Fatalf("Login request failed: %v", err)
	}
	defer response.Body.Close()

	// Step 3: Check cookies in the cookie jar for valid authentication
	isAuthenticated := false
	cookies := jar.Cookies(response.Request.URL)
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			log.Printf("Cookie 'auth_token' found, Value: %s\n", cookie.Value)
			isAuthenticated = true
		}
	}
	if !isAuthenticated {
		log.Fatalln("User failed authentication")
	}

	// Step 4: Access endpoints as authenticated user

}
