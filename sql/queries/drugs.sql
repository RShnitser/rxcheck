-- name: CreateDrug :one
INSERT INTO drugs (id, generic_name, brand_name, classification_id)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: DeleteDrugs :exec
DELETE FROM drugs;


-- name: GetDrugByGenericName :one
SELECT * FROM drugs
WHERE generic_name = ?;

-- name: ListDrugsByClassification :many
SELECT
    classifications.name as classification,
    group_concat(drugs.generic_name) as drugs
FROM 
    drugs
    JOIN classifications ON drugs.classification_id = classifications.id
GROUP BY classifications.name
ORDER BY classifications.name;