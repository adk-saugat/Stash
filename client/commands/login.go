package commands

import (
	"fmt"

	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

type LoginCommand struct{}

func (c *LoginCommand) Name() string        { return "login" }
func (c *LoginCommand) Description() string { return "Login or create an account" }

func (c *LoginCommand) Run(args []string) error {
	username, err := utils.RequireArg(args, 0, "username")
	if err != nil {
		return fmt.Errorf("%w\n\tUsage: stash login <username> <password>", err)
	}

	password, err := utils.RequireArg(args, 1, "password")
	if err != nil {
		return fmt.Errorf("%w\n\tUsage: stash login <username> <password>", err)
	}

	user := models.NewUser(1, username, password)
	user.LoginUser()
	return nil
}
