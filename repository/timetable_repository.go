package repository

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TimetableRepository struct {
	DB *pgxpool.Pool
}

func NewTimetableRepository(pool *pgxpool.Pool) *TimetableRepository {
	return &TimetableRepository{
		DB: pool,
	}
}
func (r *TimetableRepository) CreateTimeTable(
	ctx context.Context,
	t models.CreateTimetable,
) error {

	query := `
	INSERT INTO timetables
	(class_id, day, period_id, subject_id, teacher_id, type)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		t.ClassID,
		t.Day,
		t.PeriodID,
		t.SubjectID,
		t.TeacherID,
		t.Type,
	)

	return err
}
