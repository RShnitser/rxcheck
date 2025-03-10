-- name: CreateDrug :one
INSERT INTO drugs (id, generic_name, brand_name, classification_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteDrugs :exec
DELETE FROM drugs;

-- name: ListDrugsByClassification :many
SELECT * FROM drugs
WHERE classification_id = $1;