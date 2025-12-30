package models

import (
	"encoding/json"
	"log"

	"github.com/adk-saugat/stash/utils"
)

type ProjectConfig struct{
	ProjectId 		string 		`json:"projectId"`
	ProjectName 	string 		`json:"projectName"`
	TrackedFile 	[]string 	`json:"trackedFile"`
	Role 			string 		`json:"role"`
}

func ProjectConfigFromJSON(data []byte) (*ProjectConfig, error) {
	var pc ProjectConfig
	err := json.Unmarshal(data, &pc)
	if err != nil {
		return nil, err
	}
	return &pc, nil
}

func NewProjectConfig(projectName string, role string) *ProjectConfig{
	return &ProjectConfig{
		ProjectId: utils.GenerateUUID(),
		ProjectName: projectName,
		TrackedFile: make([]string, 0),
		Role: "owner",
	}
}

func (pc *ProjectConfig) AddFileToTrack(filesToTrack []string){
	for _, file := range filesToTrack {
		if !utils.FileExists(file) {
			log.Fatalf("Error: File not found: %s. Please verify the file path and ensure the file exists.", file)
		}
		// Only add if not already tracked
		if !pc.isTracked(file) {
			pc.TrackedFile = append(pc.TrackedFile, file)
		}
	}
	projectConfigJSON, err := json.MarshalIndent(pc, "", "    ")
	if err != nil {
		log.Fatal("Error: Could not convert config to JSON.")
	}

	err = utils.WriteFileData(".stash/projectConfig.json", projectConfigJSON)
	if err != nil {
		log.Fatal("Error: Could not write config to file.")
	}
}

func (pc *ProjectConfig) isTracked(file string) bool {
	for _, tracked := range pc.TrackedFile {
		if tracked == file {
			return true
		}
	}
	return false
}

func (pc *ProjectConfig) Create(){	
	projectConfigJSON, err := json.MarshalIndent(pc, "", "    ")
	if err != nil {
		log.Fatal("Error: Could not convert config to JSON.")
	}

	err = utils.WriteFileData(".stash/projectConfig.json", projectConfigJSON)
	if err != nil {
		log.Fatal("Error: Could not write config to file.")
	}
}

