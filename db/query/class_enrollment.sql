-- name: CreateClassEndrollment :one
INSERT INTO class_enrollment (
  student_id, 
  class_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetClassEnrollment :one
SELECT * FROM class_enrollment
WHERE id = $1 LIMIT 1;

-- name: ListClassEnrollment :many
SELECT * FROM class_enrollment
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateClassEnrollment :one
UPDATE class_enrollment 
SET 
  student_id = $2,
  class_id = #3
WHERE id = $1
RETURNING * 
;

-- name: DeleteClassEnrollment :exec
DELETE FROM class_enrollment
WHERE id = $1;