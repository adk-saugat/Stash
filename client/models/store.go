package models

import "time"

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