-- name: CreateMentor :one
INSERT INTO mentors (
  email, 
  name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetMentor :one
SELECT * FROM mentors
WHERE id = $1 LIMIT 1;

-- name: ListMentors :many
SELECT * FROM mentors
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateMentor :one
UPDATE mentors 
SET 
  name = $2
WHERE id = $1
RETURNING * 
;

-- name: DeleteMentor :exec
DELETE FROM mentors
WHERE id = $1;