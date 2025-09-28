-- +goose Up
CREATE TABLE drugs(
    id TEXT PRIMARY KEY NOT NULL,
    generic_name TEXT UNIQUE NOT NULL,
    brand_name TEXT UNIQUE NOT NULL,
    classification_id TEXT NOT NULL,
    FOREIGN KEY (classification_id) REFERENCES classifications(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE drugs;