package commands

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	//check if all the files are to be tracked
	if strings.ToLower(fileToTrack) == "all" {
		allFilesToTrack := watchAllFiles()
		configData.AddFileToTrack(allFilesToTrack)
	}else{
		configData.AddFileToTrack([]string{fileToTrack})
	}
}

func watchAllFiles() []string{
	allFilesToTrack := make([]string, 0)

	err := filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// Skip .stash directory
		if d.IsDir() && d.Name() == ".stash" {
			return fs.SkipDir
		}
		if !d.IsDir(){
			allFilesToTrack = append(allFilesToTrack, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error: Could not watch all the files.")
	}
	
	return allFilesToTrack
}