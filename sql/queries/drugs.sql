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

-- name: GetDrugByGenericName :one
SELECT * FROM drugs
WHERE generic_name = $1;

-- name: ListDrugsByClassification :many
SELECT
    classifications.name as classification,
    ARRAY_AGG(drugs.generic_name ORDER BY drugs.generic_name)::text[] as drugs
FROM drugs
JOIN classifications ON drugs.classification_id = classifications.id
GROUP BY classifications.name
ORDER BY classifications.name;