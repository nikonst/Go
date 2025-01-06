package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func readUsersFromFile() []User {
	jsonUsersFile, err := os.Open("data/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonUsersFile.Close()

	// Read the JSON file
	byteUersValue, err := io.ReadAll(jsonUsersFile)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of Person structs
	var users []User
	if err := json.Unmarshal(byteUersValue, &users); err != nil {
		log.Fatal(err)
	}

	return users
}

// Open the JSON file
