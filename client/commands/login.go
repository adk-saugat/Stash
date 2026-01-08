package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/adk-saugat/stash/client/api"
	"github.com/adk-saugat/stash/client/models"
	"github.com/adk-saugat/stash/client/utils"
)

type LoginCommand struct{}

func (c *LoginCommand) Name() string        { return "login" }
func (c *LoginCommand) Description() string { return "Login to your account" }

func (c *LoginCommand) Run(args []string) error {
	// Check if project exists
	if !utils.FileExists(".stash/projectConfig.json") {
		return fmt.Errorf("no stash project found. Run 'stash create' first")
	}

	configBytes, err := utils.GetFileData(utils.GetHomeDir() + "/.stashConfig")
	if err != nil {
		return fmt.Errorf("could not read config. Run 'stash config <username> <email>' first")
	}

	config, err := models.GlobalUserConfigFromJSON(configBytes)
	if err != nil {
		return fmt.Errorf("could not parse config")
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Logging in as %s\n", config.UserEmail)
	fmt.Print("Enter password: ")

	password, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("could not read password")
	}
	password = strings.TrimSpace(password)

	// Try to login
	authResp, err := api.Login(config.UserEmail, password)

	if errors.Is(err, api.ErrUserNotFound) {
		// User doesn't exist - ask to create
		fmt.Print("User not found. Create account? (y/n): ")
		answer, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("could not read input")
		}
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "y" || answer == "yes" {
			authResp, err = api.Register(config.Username, config.UserEmail, password)
			if err != nil {
				return err
			}
			// Save session
			if err := models.SaveSession(authResp.Token, authResp.Email); err != nil {
				return fmt.Errorf("could not save session: %w", err)
			}
			fmt.Println(authResp.Message)
		} else {
			fmt.Println("Login cancelled.")
		}
		return nil
	}

	if err != nil {
		return err
	}

	// Save session
	if err := models.SaveSession(authResp.Token, authResp.Email); err != nil {
		return fmt.Errorf("could not save session: %w", err)
	}

	fmt.Println(authResp.Message)
	return nil
}
