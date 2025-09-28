-- +goose Up
CREATE TABLE classifications(
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE classifications;