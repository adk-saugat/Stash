package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/adk-saugat/stash/utils"
)

type Config struct{
	ProjectId 		string 		`json:"projectId"`
	ProjectName 	string 		`json:"projectName"`
	TrackedFile 	[]string 	`json:"trackedFile"`
	Role 			string 		`json:"role"`
}

func NewConfig(projectName string, role string) *Config{

	projectIdBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal("Error: Could not generate projectId.")
	}
	projectId := strings.ToLower(strings.TrimSpace(string(projectIdBytes)))


	return &Config{
		ProjectId: string(projectId),
		ProjectName: projectName,
		TrackedFile: make([]string, 0),
		Role: "owner",
	}
}

func (config *Config) AddFileToTrack(filesToTrack []string){
	for _, file := range filesToTrack {
		if !utils.FileExists(file) {
			log.Fatalf("Error: File not found: %s. Please verify the file path and ensure the file exists.", file)
		}
		// Only add if not already tracked
		if !config.isTracked(file) {
			config.TrackedFile = append(config.TrackedFile, file)
		}
	}
	configJSON, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Fatal("Error: Couldnot convert config to JSON.")
	}

	configFile, err := os.OpenFile("./.stash/config.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error: Couldnot open config file.")
	}
	defer configFile.Close()

	_, err = configFile.Write(configJSON)
	if err != nil {
		log.Fatal("Error: Couldnot write config to file.")
	}
}

func (config *Config) isTracked(file string) bool {
	for _, tracked := range config.TrackedFile {
		if tracked == file {
			return true
		}
	}
	return false
}

func (config *Config) Create(){	
	configFile , err := os.Create(".stash/config.json")
	if err != nil {
		log.Fatal("Error: Couldnot create config file.")
	}

	configJSON, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Fatal("Error: Couldnot convert config to JSON.")
	}

	_, err = configFile.Write(configJSON)
	if err != nil {
		log.Fatal("Error: Couldnot write config to file.")
	}

	fmt.Println("Config file created!")
}