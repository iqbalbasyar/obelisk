-- name: CreateClass :one
INSERT INTO class (
  name
) VALUES (
  $1
)
RETURNING *;



-- name: GetClass :one
SELECT * FROM class
WHERE id = $1 LIMIT 1;

-- name: ListClass :many
SELECT * FROM class
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateClass :one
UPDATE class 
SET 
  name = $2
WHERE id = $1
RETURNING * 
;

-- name: DeleteClass :exec
DELETE FROM class
WHERE id = $1;