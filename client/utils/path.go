package utils

import (
	"log"
	"os"
	"strings"
)

func GetCurrentDirName() string {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error: Could not access current directory.")
	}
	currDirSlice := strings.Split(currDir, "/")
	return currDirSlice[len(currDirSlice)-1]
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error: Could not access home directory.")
	}
	return homeDir
}
