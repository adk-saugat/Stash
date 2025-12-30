package commands

import (
	"encoding/json"
	"fmt"

	emailVerifier "github.com/AfterShip/email-verifier"
	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

type ConfigCommand struct{}

func (c *ConfigCommand) Name() string {
	return "config"
}

func (c *ConfigCommand) Description() string {
	return "Configure username and email globally"
}

func (c *ConfigCommand) Run(args []string) error {
	username := utils.GetArg(1, "Error: Username not found.")
	userEmail := utils.GetArg(2, "Error: User Email not found.")

	if !emailVerifier.IsAddressValid(userEmail) {
		return fmt.Errorf("could not validate %q as email", userEmail)
	}
	fmt.Println("Email validated.")

	config := models.GlobalUserConfig{
		Username:  username,
		UserEmail: userEmail,
	}

	configData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = utils.WriteFileData(utils.GetHomeDir()+"/.stashConfig", configData)
	if err != nil {
		return err
	}

	fmt.Println("Configuration saved.")
	return nil
}