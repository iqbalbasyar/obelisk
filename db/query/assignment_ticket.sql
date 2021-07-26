-- name: CreateAssignmentTicket :one
INSERT INTO assignment_tickets (
  material_id, 
  student_id, 
  mentor_id, 
  score
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: GetAssignmentTicket :one
SELECT * FROM assignment_tickets
WHERE id = $1 LIMIT 1;

-- name: ListAssignmentTickets :many
SELECT * FROM assignment_tickets
ORDER BY id
LIMIT $1 
OFFSET $2 ;

-- name: UpdateAssignmentTicket :one
UPDATE assignment_tickets 
SET 
  material_id = $2, 
  student_id = $3,
  mentor_id = $4,
  score = $5

WHERE id = $1
RETURNING * 
;

-- name: DeleteAssignmentTicket :exec
DELETE FROM assignment_tickets
WHERE id = $1;