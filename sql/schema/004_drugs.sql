-- +goose Up
CREATE TABLE drugs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    generic_name TEXT UNIQUE NOT NULL,
    brand_name TEXT UNIQUE NOT NULL,
    classification_id UUID NOT NULL,
    FOREIGN KEY (classification_id) REFERENCES classifications(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE drugs;