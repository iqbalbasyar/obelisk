// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type AssignmentTicket struct {
	ID         int64         `json:"id"`
	MaterialID int64         `json:"material_id"`
	StudentID  int64         `json:"student_id"`
	MentorID   sql.NullInt64 `json:"mentor_id"`
	Score      sql.NullInt32 `json:"score"`
}

type Class struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ClassEnrollment struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"student_id"`
	ClassID   int64 `json:"class_id"`
}

type ClassMaterial struct {
	ID         int64 `json:"id"`
	MaterialID int64 `json:"material_id"`
	ClassID    int64 `json:"class_id"`
}

type Material struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Mentor struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	JoinedAt time.Time `json:"joined_at"`
}

type Student struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	JoinedAt time.Time `json:"joined_at"`
}
