package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"algorit.ma/iqbal/piket/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomStudent(t *testing.T) Student {
	arg := CreateStudentParams{
		Email: util.RandomEmail(),
		Name:  util.RandomName(),
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)

	// Add Requirements for testing
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, arg.Email, student.Email)
	require.Equal(t, arg.Name, student.Name)

	require.NotZero(t, student.ID)
	require.NotZero(t, student.JoinedAt)

	return student
}

func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}

func TestGetStudent(t *testing.T) {
	student1 := createRandomStudent(t)
	student2, err := testQueries.GetStudent(context.Background(), student1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, student2)
	require.Equal(t, student1.ID, student2.ID)
	require.Equal(t, student1.Email, student2.Email)
	require.Equal(t, student1.Name, student2.Name)
	require.WithinDuration(t, student1.JoinedAt, student2.JoinedAt, time.Second)

}

func TestUpdateStudent(t *testing.T) {
	student1 := createRandomStudent(t)

	arg := UpdateStudentParams{
		ID:   student1.ID,
		Name: util.RandomName(),
	}

	student2, err := testQueries.UpdateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student2)
	require.Equal(t, student1.ID, student2.ID)
	require.Equal(t, student1.Email, student2.Email)
	require.Equal(t, arg.Name, student2.Name)
	require.WithinDuration(t, student1.JoinedAt, student2.JoinedAt, time.Second)
}

func TestDeleteStudent(t *testing.T) {
	student1 := createRandomStudent(t)

	err := testQueries.DeleteStudent(context.Background(), student1.ID)
	require.NoError(t, err)

	student2, err := testQueries.GetStudent(context.Background(), student1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, student2)
}

func TestListStudents(t *testing.T) {
	// Create 10 random Account
	for i := 0; i < 10; i++ {
		createRandomStudent(t)
	}

	arg := ListStudentsParams{
		Limit:  5,
		Offset: 5,
	}

	students, err := testQueries.ListStudents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, students, 5)

	for _, student := range students {
		require.NotEmpty(t, student)

	}
}
