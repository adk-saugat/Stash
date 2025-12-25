package utils

import (
	"fmt"
	"log"
	"os"
)

// GetArg gets an argument from os.Args (legacy, for backward compatibility)
func GetArg(index int, errorMsg string) string {
	actualIndex := index + 1
	if len(os.Args) < actualIndex+1 || os.Args[actualIndex] == "" {
		log.Fatal(errorMsg)
	}
	return os.Args[actualIndex]
}

// RequireArg gets an argument from args slice or returns an error
func RequireArg(args []string, index int, name string) (string, error) {
	if len(args) < index+1 || args[index] == "" {
		return "", fmt.Errorf("%s is required", name)
	}
	return args[index], nil
}

// GetArgOrDefault gets an argument from args slice or returns a default value
func GetArgOrDefault(args []string, index int, defaultValue string) string {
	if len(args) < index+1 || args[index] == "" {
		return defaultValue
	}
	return args[index]
}
