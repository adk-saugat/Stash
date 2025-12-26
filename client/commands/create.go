package commands

import (
	"fmt"
	"os"

	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

type CreateCommand struct{}

func (c *CreateCommand) Name() string        { return "create" }
func (c *CreateCommand) Description() string { return "Create a new stash project" }

func (c *CreateCommand) Run(args []string) error {
	// creating the .stash folder
	err := utils.EnsureFolderExists(".stash")
	if err == os.ErrExist {
		return fmt.Errorf("project already exists. Cannot create a new project in an existing stash repository")
	}
	if err != nil {
		return fmt.Errorf("could not check or create folder")
	}

	// creating the store folder
	err = utils.EnsureFolderExists(".stash/store")
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("could not check or create store folder")
	}

	// check if project name added if not default to the directory name
	projectName := utils.GetArgOrDefault(args, 0, utils.GetCurrentDirName())

	projectConfig := models.NewProjectConfig(projectName, "owner")
	projectConfig.Create()
	return nil
}
