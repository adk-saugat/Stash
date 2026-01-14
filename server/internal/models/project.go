package models

import (
	"context"
	"time"

	"github.com/adk-saugat/stash/server/pkg/database"
)

type Project struct {
	ProjectId string    `json:"project_id"`
	Name      string    `json:"name"`
	OwnerId   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Project) FindById(projectId string) error {
	return database.Pool.QueryRow(context.Background(),
		"SELECT project_id, name, owner_id, created_at FROM projects WHERE project_id = $1",
		projectId,
	).Scan(&p.ProjectId, &p.Name, &p.OwnerId, &p.CreatedAt)
}

func (p *Project) Exists(projectId string) bool {
	var exists bool
	database.Pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM projects WHERE project_id = $1)",
		projectId,
	).Scan(&exists)
	return exists
}

func (p *Project) Create() error {
	_, err := database.Pool.Exec(context.Background(),
		"INSERT INTO projects (project_id, name, owner_id) VALUES ($1, $2, $3)",
		p.ProjectId, p.Name, p.OwnerId,
	)
	return err
}
