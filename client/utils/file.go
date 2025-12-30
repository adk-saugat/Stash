package utils

import (
	"errors"
	"os"
)

// FileExists checks if a file exists at the given path
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetFileData reads and returns the contents of a file
func GetFileData(path string) ([]byte, error) {
	if !FileExists(path) {
		return nil, errors.New("file does not exist: " + path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// WriteFileData writes data to a file, creating it if it doesn't exist
func WriteFileData(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
