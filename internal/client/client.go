package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xyzcv979/todo-cli/internal"
)

func sendLoginRequest(username string, password string) ([]byte, error) {
	// Prepare form data for login
	// Create a map for JSON data
	loginData := map[string]string{
		"username": username,
		"password": password,
	}

	// Marshal the map to JSON
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %v", err)
	}

	// Create the POST request with form data
	req, err := http.NewRequest("POST", internal.ServerURL+"/login", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Errorf("Error creating request: %v", err)
		return nil, err
	}

	// Set the Content-Type header for form-encoded data
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the login request to the server
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error reading response body: %v", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// If not, read the body to get the error message
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(body))
	}
	// Print the server's response (e.g., token or error message)
	fmt.Printf("Response: %s", body)
	return body, nil
}

type Token struct {
	Token string `json:"token"`
}

const TokenFile = "../../token.json"

// Write the token to a file
func writeTokenToFile(token string) error {
	t := Token{Token: token}
	file, err := os.Create(TokenFile)
	if err != nil {
		return fmt.Errorf("error creating token file: %v", err)
	}
	defer file.Close()

	// Marshal the token into JSON and write to file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(t)
	if err != nil {
		return fmt.Errorf("error writing token to file: %v", err)
	}
	return nil
}

// Read the token from the file
func readTokenFromFile() (string, error) {
	// Check if the token file exists
	if _, err := os.Stat(TokenFile); os.IsNotExist(err) {
		return "", fmt.Errorf("token file does not exist, please authenticate first")
	}

	file, err := os.Open(TokenFile)
	if err != nil {
		return "", fmt.Errorf("error opening token file: %v", err)
	}
	defer file.Close()

	var t Token
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&t)
	if err != nil {
		return "", fmt.Errorf("error reading token from file: %v", err)
	}
	return t.Token, nil
}

func main() {
	fmt.Println("client running")
	username := "admin"
	password := "password"

	body, err := sendLoginRequest(username, password)
	if err != nil {
		fmt.Errorf("%s failed authentication with error: %s", username, err)
		os.Exit(1)
	}
	fmt.Printf("%s authenticated successfully\n", username)

	var req Token
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write the token to local file for storage
	token := req.Token
	err = writeTokenToFile(token)
	if err != nil {
		fmt.Errorf("Error storing token:", err)
		os.Exit(1)
	}

	for {
		fmt.Print("Enter a command: ")
		var command string
		fmt.Scan(&command)

		token, err = readTokenFromFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
