package commands

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

type WatchCommand struct{}

func (c *WatchCommand) Name() string        { return "watch" }
func (c *WatchCommand) Description() string { return "Add files to track" }

func (c *WatchCommand) Run(args []string) error {
	fileToTrack, err := utils.RequireArg(args, 0, "file to track")
	if err != nil {
		return fmt.Errorf("%w\n\tUsage: stash watch <file|all>", err)
	}

	// read config file
	configByteData, err := utils.GetFileData("./.stash/config.json")
	if err != nil {
		return fmt.Errorf("config file not found. Run 'stash create' first")
	}

	var configData *models.Config
	err = json.Unmarshal(configByteData, &configData)
	if err != nil {
		return fmt.Errorf("could not unmarshal config data")
	}

	// check if all the files are to be tracked
	if strings.ToLower(fileToTrack) == "all" {
		allFilesToTrack := watchAllFiles()
		configData.AddFileToTrack(allFilesToTrack)
	} else {
		// Check if the path is a directory
		isDir, err := utils.FolderExists(fileToTrack)
		if err != nil {
			return fmt.Errorf("could not check path: %s", fileToTrack)
		}

		if isDir {
			// If it's a directory, walk through it to get all files
			filesToTrack := watchDirectory(fileToTrack)
			configData.AddFileToTrack(filesToTrack)
		} else {
			// Check if file exists
			if !utils.FileExists(fileToTrack) {
				return fmt.Errorf("file not found: %s", fileToTrack)
			}
			configData.AddFileToTrack([]string{fileToTrack})
		}
	}

	return nil
}

func watchAllFiles() []string {
	return watchDirectory("./")
}

func watchDirectory(dir string) []string {
	filesToTrack := make([]string, 0)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip .stash directory
		if d.IsDir() && d.Name() == ".stash" {
			return fs.SkipDir
		}
		if !d.IsDir() {
			filesToTrack = append(filesToTrack, path)
		}
		return nil
	})

	if err != nil {
		return []string{}
	}

	return filesToTrack
}
