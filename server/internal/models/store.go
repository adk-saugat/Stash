package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/adk-saugat/stash/server/pkg/database"
)

type File struct {
	Path    string `json:"path"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

type Store struct {
	StoreId   string    `json:"store_id"`
	ProjectId string    `json:"project_id"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	Date      time.Time `json:"date"`
	Files     []File    `json:"files"`
}

func (s *Store) Create() error {
	filesJSON, err := json.Marshal(s.Files)
	if err != nil {
		return err
	}

	_, err = database.Pool.Exec(context.Background(),
		"INSERT INTO stores (store_id, project_id, author, message, date, files) VALUES ($1, $2, $3, $4, $5, $6)",
		s.StoreId, s.ProjectId, s.Author, s.Message, s.Date, filesJSON,
	)
	return err
}
