-- name: CreateMaterial :one
INSERT INTO materials (
  name
) VALUES (
  $1
)
RETURNING *;



-- name: GetMaterial :one
SELECT * FROM materials
WHERE id = $1 LIMIT 1;

-- name: ListMaterials :many
SELECT * FROM materials
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateMaterial :one
UPDATE materials 
SET 
  name = $2
WHERE id = $1
RETURNING * 
;

-- name: DeleteMaterial :exec
DELETE FROM materials
WHERE id = $1;