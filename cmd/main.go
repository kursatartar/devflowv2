package main

import (
	"fmt"
	"github.com/kursatartar/devflowv2/handlers"
)

func main() {
	fmt.Println("DevFlow v0.2 - Data Management")

	handlers.CreateUser("1", "Kürşat")
	handlers.CreateUser("2", "Burak")
	handlers.ListUsers()

	handlers.UpdateUser("1", "Kürşat Artar")
	handlers.DeleteUser("2")

	handlers.ListUsers()
}
