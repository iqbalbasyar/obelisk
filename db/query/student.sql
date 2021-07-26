-- name: CreateStudent :one
INSERT INTO students (
  email, 
  name
) VALUES (
  $1, $2
)
RETURNING *;



-- name: GetStudent :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM students
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateStudent :one
UPDATE students 
SET 
  name = $2
WHERE id = $1
RETURNING * 
;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = $1;