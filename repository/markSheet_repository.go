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
	VALUES ($1,$2,$3,$4)
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

    MAX(CASE WHEN sub.name='Math' THEN m.marks END) AS math,
    MAX(CASE WHEN sub.name='Science' THEN m.marks END) AS science,
    MAX(CASE WHEN sub.name='Hindi' THEN m.marks END) AS hindi,
    MAX(CASE WHEN sub.name='English' THEN m.marks END) AS english,
    MAX(CASE WHEN sub.name='Computer' THEN m.marks END) AS computer,
    MAX(CASE WHEN sub.name='Social' THEN m.marks END) AS social,

    COALESCE(SUM(m.marks),0) AS total,

    COALESCE(ROUND((SUM(m.marks) * 100.0) / 300, 2),0) AS percentage

FROM students st
JOIN users u ON st.user_id = u.id
LEFT JOIN marks m 
ON st.id = m.student_id AND m.term = $1
LEFT JOIN subjects sub 
ON sub.id = m.subject_id

GROUP BY u.first_name, u.last_name
ORDER BY u.first_name;
	`

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

	return result, nil
}
