-- +goose Up
CREATE TABLE questions(
    id TEXT PRIMARY KEY NOT NULL,
    classification_id TEXT NOT NULL,
    drug_id TEXT NOT NULL,
    text TEXT NOT NULL,
    choice_1 TEXT NOT NULL,
    choice_2 TEXT NOT NULL,
    choice_3 TEXT NOT NULL,
    choice_4 TEXT NOT NULL,
    explanation TEXT NOT NULL,
    answer_index INT NOT NULL,
    FOREIGN KEY (classification_id) REFERENCES classifications(id) ON DELETE CASCADE,
    FOREIGN KEY (drug_id) REFERENCES drugs(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE questions;