package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

func (config *Config) Create(){
	fmt.Println("Config file created!")

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
}