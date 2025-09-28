-- name: CreateClassification :one
INSERT INTO classifications(id, name) 
VALUES (?, ?)
RETURNING *;

-- name: DeleteClassifications :exec
DELETE FROM classifications;


-- name: ListClassifications :many
SELECT * FROM classifications;


-- name: GetClassificationByName :one
SELECT * FROM classifications
WHERE name = ?;