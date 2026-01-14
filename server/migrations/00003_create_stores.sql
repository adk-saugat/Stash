-- +goose Up
CREATE TABLE stores (
    store_id UUID PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES projects(project_id),
    author VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    files JSONB NOT NULL
);

-- +goose Down
DROP TABLE stores;
