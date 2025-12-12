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

