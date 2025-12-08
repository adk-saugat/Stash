package commands

import (
	"log"
	"os"

	"github.com/adk-saugat/stash/models"
)

func Register(){
	if len(os.Args) < 3 || os.Args[2] == "" {
		log.Fatal("Error: username is required.\n\tUsage: register <username> <password>")
	}
	username := os.Args[2]

	if len(os.Args) < 4 || os.Args[3] == "" {
		log.Fatal("Error: password is required.\n\tUsage: register <username> <password>")
	}
	password := os.Args[3]

	// fmt.Printf("Username: %v\n", username)
	// fmt.Printf("Password: %v\n", password)

	user := models.NewUser(1, username, password)
	user.RegisterUser()
}