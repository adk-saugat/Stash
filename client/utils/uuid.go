package utils

import (
	"log"
	"os/exec"
	"strings"
)

func GenerateUUID() string {
	uuidBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal("Error: Could not generate UUID.")
	}
	return strings.ToLower(strings.TrimSpace(string(uuidBytes)))
}

