// Code generated by sqlc. DO NOT EDIT.
// source: class_enrollment.sql

package db

import (
	"context"
)

const createClassEndrollment = `-- name: CreateClassEndrollment :one
INSERT INTO class_enrollment (
  student_id, 
  class_id
) VALUES (
  $1, $2
)
RETURNING id, student_id, class_id
`

type CreateClassEndrollmentParams struct {
	StudentID int64 `json:"student_id"`
	ClassID   int64 `json:"class_id"`
}

func (q *Queries) CreateClassEndrollment(ctx context.Context, arg CreateClassEndrollmentParams) (ClassEnrollment, error) {
	row := q.db.QueryRowContext(ctx, createClassEndrollment, arg.StudentID, arg.ClassID)
	var i ClassEnrollment
	err := row.Scan(&i.ID, &i.StudentID, &i.ClassID)
	return i, err
}

const deleteClassEnrollment = `-- name: DeleteClassEnrollment :exec
DELETE FROM class_enrollment
WHERE id = $1
`

func (q *Queries) DeleteClassEnrollment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteClassEnrollment, id)
	return err
}

const getClassEnrollment = `-- name: GetClassEnrollment :one
SELECT id, student_id, class_id FROM class_enrollment
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetClassEnrollment(ctx context.Context, id int64) (ClassEnrollment, error) {
	row := q.db.QueryRowContext(ctx, getClassEnrollment, id)
	var i ClassEnrollment
	err := row.Scan(&i.ID, &i.StudentID, &i.ClassID)
	return i, err
}

const listClassEnrollment = `-- name: ListClassEnrollment :many
SELECT id, student_id, class_id FROM class_enrollment
ORDER BY id
LIMIT $1 
OFFSET $2
`

type ListClassEnrollmentParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListClassEnrollment(ctx context.Context, arg ListClassEnrollmentParams) ([]ClassEnrollment, error) {
	rows, err := q.db.QueryContext(ctx, listClassEnrollment, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClassEnrollment
	for rows.Next() {
		var i ClassEnrollment
		if err := rows.Scan(&i.ID, &i.StudentID, &i.ClassID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClassEnrollment = `-- name: UpdateClassEnrollment :one
UPDATE class_enrollment 
SET 
  student_id = $2,
  class_id = #3
WHERE id = $1
RETURNING id, student_id, class_id
`

type UpdateClassEnrollmentParams struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"student_id"`
}

func (q *Queries) UpdateClassEnrollment(ctx context.Context, arg UpdateClassEnrollmentParams) (ClassEnrollment, error) {
	row := q.db.QueryRowContext(ctx, updateClassEnrollment, arg.ID, arg.StudentID)
	var i ClassEnrollment
	err := row.Scan(&i.ID, &i.StudentID, &i.ClassID)
	return i, err
}