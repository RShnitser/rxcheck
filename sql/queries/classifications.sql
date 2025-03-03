-- name: CreateClassification :one
INSERT INTO classifications (id, name)
VALUES (
    gen_random_uuid(),
    $1
)
RETURNING *;

-- name: DeleteClassifications :exec
DELETE FROM classifications;

-- name: GetClassificationByUserName :one
SELECT * FROM classifications
WHERE name = $1;