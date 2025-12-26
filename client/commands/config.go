package commands

import (
	"encoding/json"
	"fmt"
	"os"

	emailVerifier "github.com/AfterShip/email-verifier"
	"github.com/adk-saugat/stash/utils"
)

type ConfigCommand struct{}

type configUser struct {
	Username  string `json:"username"`
	UserEmail string `json:"userEmail"`
}

func (c *ConfigCommand) Name() string {
	 return "config" 
}

func (c *ConfigCommand) Description() string { 
	return "Configure username and email globally" 
}

func (c *ConfigCommand) Run(args []string) error {
	username := utils.GetArg(1, "Error: Username not found.")
	userEmail := utils.GetArg(2, "Error: User Email not found.")

	isValid := emailVerifier.IsAddressValid(userEmail)
	if !isValid{
		return fmt.Errorf("Could not validate %v as email.", userEmail)
	}
	homeDir, _ := os.UserHomeDir()

	stashConfig, err := os.Create(homeDir + "/.stashConfig")
	if err != nil {
		return err
	}

	config := configUser{
		Username:  username,
		UserEmail: userEmail,
	}

	configData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	if _, err = stashConfig.WriteString(string(configData));err != nil {
		return err
	}

	return nil
}