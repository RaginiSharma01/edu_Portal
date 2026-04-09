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

	var subjectID interface{}
	var teacherID interface{}

	// Treat empty string as NULL
	if t.SubjectID != nil && *t.SubjectID != "" {
		subjectID = *t.SubjectID
	}

	if t.TeacherID != nil && *t.TeacherID != "" {
		teacherID = *t.TeacherID
	}

	_, err := r.DB.Exec(
		ctx,
		query,
		t.ClassID,
		t.Day,
		t.PeriodID,
		subjectID,
		teacherID,
		t.Type,
	)

	return err
}

func (r *TimetableRepository) GetTimetable(
	ctx context.Context,
	classID string,
) ([]models.TimetableRow, error) {

	query := `
	SELECT 
		p.period_number,
		p.start_time::text,
		p.end_time::text,

		MAX(CASE WHEN t.day = 'Monday' THEN COALESCE(s.name, t.type) END) AS monday,
		MAX(CASE WHEN t.day = 'Tuesday' THEN COALESCE(s.name, t.type) END) AS tuesday,
		MAX(CASE WHEN t.day = 'Wednesday' THEN COALESCE(s.name, t.type) END) AS wednesday,
		MAX(CASE WHEN t.day = 'Thursday' THEN COALESCE(s.name, t.type) END) AS thursday,
		MAX(CASE WHEN t.day = 'Friday' THEN COALESCE(s.name, t.type) END) AS friday

	FROM periods p

	LEFT JOIN timetables t
		ON p.id = t.period_id
		AND t.class_id = $1

	LEFT JOIN subjects s
		ON s.id = t.subject_id

	GROUP BY p.id, p.period_number, p.start_time, p.end_time
	ORDER BY p.period_number;
	`

	rows, err := r.DB.Query(ctx, query, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []models.TimetableRow

	for rows.Next() {
		var row models.TimetableRow

		err := rows.Scan(
			&row.PeriodNumber,
			&row.StartTime,
			&row.EndTime,
			&row.Monday,
			&row.Tuesday,
			&row.Wednesday,
			&row.Thursday,
			&row.Friday,
		)
		if err != nil {
			return nil, err
		}

		timetable = append(timetable, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return timetable, nil
}