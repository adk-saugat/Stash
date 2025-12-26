package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/adk-saugat/stash/models"
	"github.com/adk-saugat/stash/utils"
)

type StoreCommand struct{}

func (c *StoreCommand) Name() string        { return "store" }
func (c *StoreCommand) Description() string { return "Store current changes" }

func (c *StoreCommand) Run(args []string) error {
	_, err := utils.RequireArg(args, 0, "store message")
	if err != nil {
		return fmt.Errorf("%w\n\tUsage: stash store <message>", err)
	}

	// storeMessage := args[0]

	projectConfigBytes, err := utils.GetFileData(".stash/projectConfig.json")
	if err != nil {
		return fmt.Errorf("could not read config. Run 'stash create' first")
	}

	projectConfig, err := models.ProjectConfigFromJSON(projectConfigBytes)
	if err != nil {
		return fmt.Errorf("could not parse config")
	}

	storeFiles := make([]models.File, 0)
	for _, filePath := range projectConfig.TrackedFile {
		fileData, err := utils.GetFileData(filePath)
		if err != nil {
			return fmt.Errorf("could not read file: %s", filePath)
		}

		fileHash := utils.GenerateSHA256(fileData)
		fileContent := base64.StdEncoding.EncodeToString(fileData)
		storeFile := models.NewFile(filePath, fileHash, fileContent)
		storeFiles = append(storeFiles, storeFile)
	}

	fmt.Printf("Stored %d files\n", len(storeFiles))

	// TODO: Create a store struct and add the store to the .stash.
	
	return nil
}
