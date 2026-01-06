package commands

import (
	"fmt"
	"os"

	"github.com/adk-saugat/stash/client/models"
	"github.com/adk-saugat/stash/client/utils"
)

type CreateCommand struct{}

func (c *CreateCommand) Name() string        { return "create" }
func (c *CreateCommand) Description() string { return "Create a new stash project" }

func (c *CreateCommand) Run(args []string) error {
	err := utils.EnsureFolderExists(".stash")
	if err == os.ErrExist {
		return fmt.Errorf("project already exists. Cannot create a new project in an existing stash repository")
	}
	if err != nil {
		return fmt.Errorf("could not check or create folder")
	}
	fmt.Println("Repository initialized.")

	projectName := utils.GetArgOrDefault(args, 0, utils.GetCurrentDirName())

	projectConfig := models.NewProjectConfig(projectName, "owner")
	projectConfig.Create()
	fmt.Println("Project configuration created.")

	return nil
}
