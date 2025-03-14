-- +goose Up
CREATE TABLE answers(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text TEXT NOT NULL,
    answer_order INT NOT NULL DEFAULT 0,
    question_id UUID NOT NULL,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE answers;