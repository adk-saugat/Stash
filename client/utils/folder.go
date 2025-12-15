package utils

import (
	"os"
)

// EnsureFolderExists checks if folder exists and returns an error if it does.
// If folder doesn't exist, it creates the folder.
func EnsureFolderExists(path string) error {
	exists, err := FolderExists(path)
	if err != nil {
		return err
	}

	if exists {
		return os.ErrExist
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

// checks if a folder exists at the given path
func FolderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

