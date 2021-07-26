-- name: CreateClassMaterial :one
INSERT INTO class_materials (
  material_id, 
  class_id
) VALUES (
  $1, $2
)
RETURNING *;



-- name: GetClassMaterial :one
SELECT * FROM class_materials
WHERE id = $1 LIMIT 1;

-- name: ListClassMaterials :many
SELECT * FROM class_materials
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateClassMaterial :one
UPDATE class_materials 
SET 
  material_id = $2,
  class_id = $3
WHERE id = $1
RETURNING * 
;

-- name: DeleteClassMaterial :exec
DELETE FROM class_materials
WHERE id = $1;