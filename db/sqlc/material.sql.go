// Code generated by sqlc. DO NOT EDIT.
// source: material.sql

package db

import (
	"context"
)

const createMaterial = `-- name: CreateMaterial :one
INSERT INTO materials (
  name
) VALUES (
  $1
)
RETURNING id, name
`

func (q *Queries) CreateMaterial(ctx context.Context, name string) (Material, error) {
	row := q.db.QueryRowContext(ctx, createMaterial, name)
	var i Material
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteMaterial = `-- name: DeleteMaterial :exec
DELETE FROM materials
WHERE id = $1
`

func (q *Queries) DeleteMaterial(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMaterial, id)
	return err
}

const getMaterial = `-- name: GetMaterial :one
SELECT id, name FROM materials
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMaterial(ctx context.Context, id int64) (Material, error) {
	row := q.db.QueryRowContext(ctx, getMaterial, id)
	var i Material
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listMaterials = `-- name: ListMaterials :many
SELECT id, name FROM materials
ORDER BY id
LIMIT $1 
OFFSET $2
`

type ListMaterialsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMaterials(ctx context.Context, arg ListMaterialsParams) ([]Material, error) {
	rows, err := q.db.QueryContext(ctx, listMaterials, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Material
	for rows.Next() {
		var i Material
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const updateMaterial = `-- name: UpdateMaterial :one
UPDATE materials 
SET 
  name = $2
WHERE id = $1
RETURNING id, name
`

type UpdateMaterialParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateMaterial(ctx context.Context, arg UpdateMaterialParams) (Material, error) {
	row := q.db.QueryRowContext(ctx, updateMaterial, arg.ID, arg.Name)
	var i Material
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
