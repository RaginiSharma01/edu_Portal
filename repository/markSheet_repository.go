package repository

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MarksRepository struct {
	DB *pgxpool.Pool
}

func NewMarksRepository(pool *pgxpool.Pool) *MarksRepository {
	return &MarksRepository{
		DB: pool,
	}
}

func (r *MarksRepository) CreateMarks(
	ctx context.Context,
	m models.CreateMarks,
) error {

	query := `
	INSERT INTO marks (student_id, subject_id, term, marks)
	VALUES ($1, $2, $3, $4)
	`

	for _, sub := range m.Marks {

		_, err := r.DB.Exec(
			ctx,
			query,
			m.StudentID,
			sub.SubjectID,
			m.Term,
			sub.Marks,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *MarksRepository) GetMarks(
	ctx context.Context,
	term string,
) ([]models.StudentMarks, error) {

	query := `
SELECT 
    u.first_name || ' ' || u.last_name AS student,

 COALESCE(MAX(CASE WHEN LOWER(sub.name) = 'math' THEN m.marks END),0) AS math,
    COALESCE(MAX(CASE WHEN LOWER(sub.name) = 'science' THEN m.marks END),0) AS science,
    COALESCE(MAX(CASE WHEN LOWER(sub.name) = 'hindi' THEN m.marks END),0) AS hindi,
    COALESCE(MAX(CASE WHEN LOWER(sub.name) = 'english' THEN m.marks END),0) AS english,
    COALESCE(MAX(CASE WHEN LOWER(sub.name) = 'computer' THEN m.marks END),0) AS computer,
    COALESCE(MAX(CASE WHEN LOWER(sub.name) = 'social' THEN m.marks END),0) AS social,

    COALESCE(SUM(m.marks),0) AS total,
    COALESCE(ROUND((SUM(m.marks) * 100.0) / 300, 2),0) AS percentage

FROM users u

LEFT JOIN marks m
ON u.id = m.student_id AND (m.term = $1 OR m.term = 'Term'||$1)

LEFT JOIN subjects sub
ON sub.id = m.subject_id

WHERE u.role = 'student'

GROUP BY u.id, u.first_name, u.last_name
ORDER BY u.first_name;`

	rows, err := r.DB.Query(ctx, query, term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.StudentMarks

	for rows.Next() {

		var s models.StudentMarks

		err := rows.Scan(
			&s.Student,
			&s.Math,
			&s.Science,
			&s.Hindi,
			&s.English,
			&s.Computer,
			&s.Social,
			&s.Total,
			&s.Percentage,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
