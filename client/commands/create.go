package commands

import (
	"log"
	"os"
	"strings"

	"github.com/adk-saugat/stash/models"
)

func Create(){
	// creating the .stash folder
	exists, err := folderExists(".stash")
	if err != nil {
		log.Fatal("Error: Couldnot check folders existence.")
	}

	if !exists {
		err = os.Mkdir(".stash", 0755)
		if err != nil {
			log.Fatal("Error: Couldnot create folder.")
		}
	}

	// creating the store folder
	exists, err = folderExists(".stash/store")
	if err != nil {
		log.Fatal("Error: Couldnot check folders existence.")
	}

	if !exists {
		err = os.Mkdir("./.stash/store", 0755)
		if err != nil {
			log.Fatal("Error: Couldnot create.")
		}
	}

	// check if project name added if not default to the directory name
	if len(os.Args) < 3 || os.Args[2] == "" {
		currDir, err := os.Getwd()
		if err != nil {
			log.Fatal("Error: Could not access current directory.")
		}
		currDirSlice := strings.Split(currDir, "/")
		projectName := currDirSlice[len(currDirSlice)-1]
		config := models.NewConfig(projectName, "owner")
		config.Create()
		return
	}

	// fmt.Println("project-name added")
	config := models.NewConfig(os.Args[2], "owner")
	config.Create()
}

func folderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil 
	}
	if os.IsNotExist(err) {
		return false, nil 
	}
	return false, err 
}
