-- +goose Up
CREATE TABLE questions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text TEXT NOT NULL,
    classification_id UUID NOT NULL,
    drug_id UUID NOT NULL,
    FOREIGN KEY (classification_id) REFERENCES classifications(id) ON DELETE CASCADE,
    FOREIGN KEY (drug_id) REFERENCES drugs(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE questions;