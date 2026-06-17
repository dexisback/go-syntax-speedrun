package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// first make a struct or schema
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func saveUser(user User) error {
	// read the json file
	fileData, err := os.ReadFile("user.json")
	if err != nil {
		return err
	}

	// convert JSON to go struct slice
	var users []User
	err = json.Unmarshal(fileData, &users)
	if err != nil {
		return err
	}

	// append new user to slice
	users = append(users, user)

	// convert back to json
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	// write entire file again
	err = os.WriteFile("user.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	user := User{
		ID:    6,
		Name:  "Vikram Patel",
		Email: "vikram@example.com",
	}

	err := saveUser(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User saved successfully")
}
