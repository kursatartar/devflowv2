package handlers

import (
	"fmt"
	"github.com/kursatartar/devflowv2/models"
)

func CreateUser(id string, name string) {
	models.Users[id] = name
	fmt.Println("User created:", name)
}

func ListUsers() {
	fmt.Println("All users:")
	for id, name := range models.Users {
		fmt.Printf("- %s: %s\n", id, name)
	}
}

func UpdateUser(id string, newName string) {
	if _, exists := models.Users[id]; exists {
		models.Users[id] = newName
		fmt.Println("User updated:", newName)
	} else {
		fmt.Println("User not found")
	}
}

func DeleteUser(id string) {
	if _, exists := models.Users[id]; exists {
		delete(models.Users, id)
		fmt.Println("User deleted:", id)
	} else {
		fmt.Println("User not found")
	}
}
