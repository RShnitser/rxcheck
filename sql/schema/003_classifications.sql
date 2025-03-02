-- +goose Up
CREATE TABLE classifications(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE classifications;
