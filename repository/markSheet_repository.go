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
	return &MarksRepository{DB: pool}
}

func (r *MarksRepository) CreateMarks(ctx context.Context, m models.CreateMarks) error {
	query := `
		INSERT INTO marks (student_id, subject_id, term, marks)
		VALUES ($1, $2, $3, $4)
	`
	for _, sub := range m.Marks {
		_, err := r.DB.Exec(ctx, query, m.StudentID, sub.SubjectID, m.Term, sub.Marks)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MarksRepository) GetMaxMarksByTerm(ctx context.Context, term string) (int, error) {
	query := `
        SELECT max_marks FROM term_config
        WHERE term = $1 OR term = 'Term' || $1
        LIMIT 1
    `
	var maxMarks int
	err := r.DB.QueryRow(ctx, query, term).Scan(&maxMarks)
	if err != nil {
		return 100, err
	}
	return maxMarks, nil
}

func (r *MarksRepository) GetMarks(ctx context.Context, term string) ([]models.StudentMarks, error) {

	maxMarks, err := r.GetMaxMarksByTerm(ctx, term)
	if err != nil {
		return nil, err
	}

	query := `
SELECT
    u.first_name || ' ' || u.last_name AS student,
    sub.name                            AS subject,
    COALESCE(m.marks, 0)               AS marks
FROM users u
CROSS JOIN subjects sub
LEFT JOIN marks m
    ON m.student_id = u.id
    AND m.subject_id = sub.id
    AND (m.term = $1 OR m.term = 'Term' || $1)
WHERE u.role = 'student'
ORDER BY u.first_name, sub.name`

	rows, err := r.DB.Query(ctx, query, term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	studentMap := make(map[string]*models.StudentMarks)
	var studentOrder []string

	for rows.Next() {
		var student, subject string
		var marks int

		if err := rows.Scan(&student, &subject, &marks); err != nil {
			return nil, err
		}

		if _, exists := studentMap[student]; !exists {
			studentMap[student] = &models.StudentMarks{
				Student:  student,
				Subjects: make(map[string]int),
			}
			studentOrder = append(studentOrder, student)
		}

		//  KEEP ORIGINAL MARKS
		studentMap[student].Subjects[subject] = marks
		studentMap[student].Total += marks
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	result := make([]models.StudentMarks, 0, len(studentOrder))

	for _, name := range studentOrder {
		s := studentMap[name]

		totalSubjects := len(s.Subjects)
		if totalSubjects == 0 {
			continue
		}

		//  FIXED MAX TOTAL LOGIC
		baseMaxTotal := 100 * totalSubjects

		if maxMarks != 100 {
			s.MaxTotal = baseMaxTotal / 2
		} else {
			s.MaxTotal = baseMaxTotal
		}

		if s.MaxTotal > 0 {
			s.Percentage = (float64(s.Total) * 100) / float64(s.MaxTotal)
		}

		result = append(result, *s)
	}

	return result, nil
}

// Get single student marks
func (r *MarksRepository) GetMarksByStudentID(
	ctx context.Context,
	studentID string,
	term string,
) (models.StudentMarks, error) {

	maxMarks, err := r.GetMaxMarksByTerm(ctx, term)
	if err != nil {
		return models.StudentMarks{}, err
	}

	query := `
SELECT
    u.first_name || ' ' || u.last_name AS student,
    sub.name                            AS subject,
    COALESCE(m.marks, 0)               AS marks
FROM users u
CROSS JOIN subjects sub
LEFT JOIN marks m
    ON m.student_id = u.id
    AND m.subject_id = sub.id
    AND (m.term = $2 OR m.term = 'Term' || $2)
WHERE u.id = $1
ORDER BY sub.name`

	rows, err := r.DB.Query(ctx, query, studentID, term)
	if err != nil {
		return models.StudentMarks{}, err
	}
	defer rows.Close()

	s := models.StudentMarks{
		Subjects: make(map[string]int),
	}

	for rows.Next() {
		var subject string
		var marks int

		if err := rows.Scan(&s.Student, &subject, &marks); err != nil {
			return models.StudentMarks{}, err
		}

		s.Subjects[subject] = marks
		s.Total += marks
	}

	if err := rows.Err(); err != nil {
		return models.StudentMarks{}, err
	}

	totalSubjects := len(s.Subjects)

	baseMaxTotal := 100 * totalSubjects

	if maxMarks == 50 {
		s.MaxTotal = baseMaxTotal / 2
	} else {
		s.MaxTotal = baseMaxTotal
	}

	//  CORRECT PERCENTAGE
	if s.MaxTotal > 0 {
		s.Percentage = (float64(s.Total) * 100) / float64(s.MaxTotal)
	}

	return s, nil
}
