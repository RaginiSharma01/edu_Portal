package repository

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClassroomRepo struct {
	db *pgxpool.Pool
}

func NewClassroomRepo(db *pgxpool.Pool) *ClassroomRepo {
	return &ClassroomRepo{db: db}
}

func (r *ClassroomRepo) CreateClassroom(ctx context.Context, req models.CreateClassroom) (string, error) {

	query := `
	INSERT INTO classrooms (name, teacher_id, academic_year)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	var classroomID string

	err := r.db.QueryRow(ctx, query,
		req.Name,
		req.TeacherID,
		req.AcademicYear,
	).Scan(&classroomID)

	if err != nil {
		return "", err
	}

	return classroomID, nil
}
