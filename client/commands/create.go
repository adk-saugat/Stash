package commands

import (
	"log"
	"os"

	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

func Create() {
	// creating the .stash folder
	err := utils.EnsureFolderExists(".stash")
	if err != nil {
		log.Fatal("Error: Couldnot check or create folder.")
	}

	// creating the store folder
	err = utils.EnsureFolderExists(".stash/store")
	if err != nil {
		log.Fatal("Error: Couldnot check or create folder.")
	}

	// check if project name added if not default to the directory name
	var projectName string
	if len(os.Args) >= 3 && os.Args[2] != "" {
		projectName = os.Args[2]
	} else {
		projectName = utils.GetCurrentDirName()
	}

	config := models.NewConfig(projectName, "owner")
	config.Create()
}
