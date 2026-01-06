package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adk-saugat/stash/client/models"
	"github.com/adk-saugat/stash/client/utils"
)

type LoginCommand struct{}

func (c *LoginCommand) Name() string        { return "login" }
func (c *LoginCommand) Description() string { return "Login to your account" }

func (c *LoginCommand) Run(args []string) error {
	configBytes, err := utils.GetFileData(utils.GetHomeDir() + "/.stashConfig")
	if err != nil {
		return fmt.Errorf("could not read config. Run 'stash config <username> <email>' first")
	}

	config, err := models.GlobalUserConfigFromJSON(configBytes)
	if err != nil {
		return fmt.Errorf("could not parse config")
	}

	fmt.Printf("Logging in as %s\n", config.UserEmail)
	fmt.Print("Enter password: ")

	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("could not read password")
	}
	password = strings.TrimSpace(password)

	user := models.NewUser(config.Username, config.UserEmail, password)
	user.LoginUser()
	return nil
}
