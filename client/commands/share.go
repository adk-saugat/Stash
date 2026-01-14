package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adk-saugat/stash/client/api"
	"github.com/adk-saugat/stash/client/models"
	"github.com/adk-saugat/stash/client/utils"
)

type ShareCommand struct{}

func (c *ShareCommand) Name() string {
	return "share"
}

func (c *ShareCommand) Description() string {
	return "Share local stores to the server"
}

func (c *ShareCommand) Run(args []string) error {
	// Check if project exists
	exists, _ := utils.FolderExists(".stash")
	if !exists {
		return fmt.Errorf("no stash project found. Run 'stash create' first")
	}

	// Check if logged in
	if !models.IsSessionValid() {
		return fmt.Errorf("not logged in. Run 'stash login' first")
	}

	session, err := models.LoadSession()
	if err != nil {
		return fmt.Errorf("could not load session: %w", err)
	}

	// Load project config
	configData, err := os.ReadFile(".stash/projectConfig.json")
	if err != nil {
		return fmt.Errorf("could not read project config: %w", err)
	}

	projectConfig, err := models.ProjectConfigFromJSON(configData)
	if err != nil {
		return fmt.Errorf("could not parse project config: %w", err)
	}

	// Check if any files are being tracked
	if len(projectConfig.TrackedFile) == 0 {
		return fmt.Errorf("no files being tracked. Run 'stash watch <files>' first")
	}

	// Get all local stores
	allStores, err := models.GetAllStores()
	if err != nil {
		return fmt.Errorf("could not get stores: %w", err)
	}

	if len(allStores) == 0 {
		return fmt.Errorf("no stores to share. Run 'stash store' first")
	}

	// Load shared store IDs
	sharedIds := loadSharedIds()

	// Find unshared stores
	var unsharedStores []models.Store
	for _, store := range allStores {
		if !sharedIds[store.StoreId] {
			unsharedStores = append(unsharedStores, store)
		}
	}

	if len(unsharedStores) == 0 {
		fmt.Println("All stores already shared.")
		return nil
	}

	fmt.Printf("Sharing %d store(s)...\n", len(unsharedStores))

	// Share each unshared store (oldest first)
	for i := len(unsharedStores) - 1; i >= 0; i-- {
		store := unsharedStores[i]
		fmt.Printf("  ● %s - %s\n", store.StoreId[:8], store.Message)

		_, err := api.ShareStore(
			session.Token,
			projectConfig.ProjectId,
			projectConfig.ProjectName,
			store,
		)
		if err != nil {
			return fmt.Errorf("failed to share store %s: %w", store.StoreId[:8], err)
		}

		// Mark as shared
		sharedIds[store.StoreId] = true
		saveSharedIds(sharedIds)
	}

	fmt.Println("All stores shared successfully!")
	return nil
}

func loadSharedIds() map[string]bool {
	sharedIds := make(map[string]bool)
	data, err := os.ReadFile(".stash/shared.json")
	if err != nil {
		return sharedIds
	}
	var ids []string
	json.Unmarshal(data, &ids)
	for _, id := range ids {
		sharedIds[id] = true
	}
	return sharedIds
}

func saveSharedIds(sharedIds map[string]bool) {
	var ids []string
	for id := range sharedIds {
		ids = append(ids, id)
	}
	data, _ := json.MarshalIndent(ids, "", "    ")
	utils.WriteFileData(".stash/shared.json", data)
}
