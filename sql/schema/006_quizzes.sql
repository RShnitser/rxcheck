-- +goose Up
CREATE TABLE quizzes(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL,
    question_1 UUID NOT NULL,
    question_2 UUID NOT NULL,
    question_3 UUID NOT NULL,
    question_4 UUID NOT NULL,
    question_5 UUID NOT NULL,
    next_question_index INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (question_1) REFERENCES questions(id),
    FOREIGN KEY (question_2) REFERENCES questions(id),
    FOREIGN KEY (question_3) REFERENCES questions(id),
    FOREIGN KEY (question_4) REFERENCES questions(id),
    FOREIGN KEY (question_5) REFERENCES questions(id)
);

-- +goose Down
DROP TABLE quizzes;