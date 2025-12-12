package commands

import (
	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

func Login() {
	username := utils.GetArg(1, "Error: username is required.\n\tUsage: login <username> <password>")
	password := utils.GetArg(2, "Error: password is required.\n\tUsage: login <username> <password>")

	user := models.NewUser(1, username, password)
	user.LoginUser()
}