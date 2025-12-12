package utils

import (
	"os"
)

// ensure folder exists and if not then creates the folder
func EnsureFolderExists(path string) error {
	exists, err := folderExists(path)
	if err != nil {
		return err
	}

	if !exists {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// checks if a folder exists at the given path
func folderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

