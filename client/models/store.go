package models

import (
	"encoding/json"
	"os"
	"time"

	"github.com/adk-saugat/stash/utils"
)

type File struct {
	Path    string `json:"path"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

func NewFile(path, hash, content string) File{
	return File{
		Path: path,
		Hash: hash,
		Content: content,
	}
}

type Store struct {
	StoreId   string    `json:"store_id"`
	ProjectId string    `json:"project_id"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	Date      time.Time `json:"date"`
	Files     []File    `json:"files"`
}

func NewStore(projectId, author, message string, files []File) Store {
	return Store{
		StoreId:   utils.GenerateUUID(),
		ProjectId: projectId,
		Author:    author,
		Message:   message,
		Date:      time.Now(),
		Files:     files,
	}
}

func (store *Store) ToJSON() ([]byte, error) {
	return json.MarshalIndent(store, "", "    ")
}

func (store *Store) Create() error {
	storeJSON, err := store.ToJSON()
	if err != nil {
		return err
	}

	// Ensure stores directory exists
	if err := os.MkdirAll("./.stash/stores", 0755); err != nil {
		return err
	}

	return utils.WriteFileData("./.stash/stores/"+store.StoreId+".json", storeJSON)
}