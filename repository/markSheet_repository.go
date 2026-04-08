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
	    st.name AS student,

	    MAX(CASE WHEN sub.name='Mathematics' THEN m.marks END) AS mathematics,
	    MAX(CASE WHEN sub.name='Physics' THEN m.marks END) AS physics,
	    MAX(CASE WHEN sub.name='Chemistry' THEN m.marks END) AS chemistry,
	    MAX(CASE WHEN sub.name='English' THEN m.marks END) AS english,
	    MAX(CASE WHEN sub.name='Hindi' THEN m.marks END) AS hindi,
	    MAX(CASE WHEN sub.name='Social Science' THEN m.marks END) AS social,

	    SUM(m.marks) AS total,
	    ROUND((SUM(m.marks) * 100.0) / 300,2) AS percentage

	FROM students st
	LEFT JOIN marks m ON st.id = m.student_id
	LEFT JOIN subjects sub ON sub.id = m.subject_id

	WHERE m.term = $1

	GROUP BY st.name
	ORDER BY st.name
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
			&s.Mathematics,
			&s.Physics,
			&s.Chemistry,
			&s.English,
			&s.Hindi,
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