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
SELECT sqlc.embed(drugs), classifications.name
FROM drugs
JOIN classifications ON drugs.classification_id = classifications.id
ORDER BY classifications.name;