-- +goose Up
CREATE TABLE sessions(
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT UNIQUE NOT NULL,
    question_1 TEXT NOT NULL,
    question_2 TEXT NOT NULL,
    question_3 TEXT NOT NULL,
    question_4 TEXT NOT NULL,
    question_5 TEXT NOT NULL,
    score INT NOT NULL,
    question_index INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (question_1) REFERENCES questions(id),
    FOREIGN KEY (question_2) REFERENCES questions(id),
    FOREIGN KEY (question_3) REFERENCES questions(id),
    FOREIGN KEY (question_4) REFERENCES questions(id),
    FOREIGN KEY (question_5) REFERENCES questions(id)
);

-- +goose Down
DROP TABLE sessions;