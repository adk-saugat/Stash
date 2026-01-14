-- +goose Up
CREATE TABLE projects (
    project_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE projects;
