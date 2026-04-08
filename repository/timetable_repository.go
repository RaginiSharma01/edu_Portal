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
	( class_id, day, period_id, subject_id, teacher_id, type)
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
func (r *TimetableRepository) GetTimetable(
	ctx context.Context,
	classID string,
) ([]models.TimetableRow, error) {

	query := `
	SELECT 
	    p.period_number,
	    p.start_time,
	    p.end_time,

	    MAX(CASE WHEN t.day = 'Monday' THEN s.name END) AS monday,
	    MAX(CASE WHEN t.day = 'Tuesday' THEN s.name END) AS tuesday,
	    MAX(CASE WHEN t.day = 'Wednesday' THEN s.name END) AS wednesday,
	    MAX(CASE WHEN t.day = 'Thursday' THEN s.name END) AS thursday,
	    MAX(CASE WHEN t.day = 'Friday' THEN s.name END) AS friday

	FROM periods p
	LEFT JOIN timetables t ON p.id = t.period_id
	LEFT JOIN subjects s ON s.id = t.subject_id
	WHERE t.class_id = $1
	GROUP BY p.period_number, p.start_time, p.end_time
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

	return timetable, nil
}
