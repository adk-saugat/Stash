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
	storeMessage, err := utils.RequireArg(args, 0, "store message")
	if err != nil {
		return fmt.Errorf("%w\n\tUsage: stash store <message>", err)
	}

	projectConfigBytes, err := utils.GetFileData(".stash/projectConfig.json")
	if err != nil {
		return fmt.Errorf("could not read config. Run 'stash create' first")
	}

	projectConfig, err := models.ProjectConfigFromJSON(projectConfigBytes)
	if err != nil {
		return fmt.Errorf("could not parse config")
	}
	fmt.Println("Project configuration loaded.")

	configByte, err := utils.GetFileData(utils.GetHomeDir() + "/.stashConfig")
	if err != nil {
		return fmt.Errorf("could not read user config. Run 'stash config <username> <email>' first")
	}

	config, err := models.GlobalUserConfigFromJSON(configByte)
	if err != nil {
		return fmt.Errorf("could not parse user config")
	}
	fmt.Println("User configuration loaded.")

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
	fmt.Println("Tracked files processed.")

	latestStore, _ := models.GetLatestStore()
	if latestStore != nil && !hasChanges(storeFiles, latestStore.Files) {
		return fmt.Errorf("no changes to store")
	}

	storeData := models.NewStore(projectConfig.ProjectId, config.UserEmail, storeMessage, storeFiles)

	err = storeData.Create()
	if err != nil {
		return fmt.Errorf("could not create store: %w", err)
	}

	fmt.Println("Store created.")
	return nil
}

func hasChanges(storeFiles []models.File, latestStoreFiles []models.File) bool {
	// Different number of files means changes
	if len(storeFiles) != len(latestStoreFiles) {
		return true
	}

	// Build hash map from latest store
	lastHashes := make(map[string]string)
	for _, f := range latestStoreFiles {
		lastHashes[f.Path] = f.Hash
	}

	// Check if any file hash changed
	for _, f := range storeFiles {
		if lastHashes[f.Path] != f.Hash {
			return true
		}
	}

	return false
}
