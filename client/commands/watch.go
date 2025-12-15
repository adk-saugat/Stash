package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

func Watch(){
	// check the file to track
	fileToTrack := utils.GetArg(1, "Error: File to track not found.")

	// read config file
	configByteData, err := os.ReadFile("./.stash/config.json")
	if err != nil {
		log.Fatal("Error: Config file not found.")
	}
	
	var configData *models.Config
	err = json.Unmarshal(configByteData, &configData)
	if err != nil {
		log.Fatal("Error: Could not unmarshal config data.")
	}
	
	configData.AddFileToTrack(fileToTrack)

}