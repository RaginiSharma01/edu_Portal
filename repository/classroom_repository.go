package repository

import (
	"context"
	"errors"
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

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	var classroomID string

	err = tx.QueryRow(ctx, `
		INSERT INTO classrooms (name, teacher_id, academic_year)
		VALUES ($1, $2, $3)
		RETURNING id
	`,
		req.Name,
		req.TeacherID,
		req.AcademicYear,
	).Scan(&classroomID)

	if err != nil {
		return "", err
	}

	for _, subjectName := range req.Subjects {

		var subjectID string

		err := tx.QueryRow(ctx,
			`SELECT id FROM subjects WHERE LOWER(name) = LOWER($1)`,
			subjectName,
		).Scan(&subjectID)

		if err != nil {
			return "", errors.New("subject not found: " + subjectName)
		}

		_, err = tx.Exec(ctx,
			`INSERT INTO classroom_subjects (classroom_id, subject_id) VALUES ($1, $2)`,
			classroomID,
			subjectID,
		)

		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}

	return classroomID, nil
}

func (r *ClassroomRepo) GetClassrooms(ctx context.Context) ([]models.ClassroomCard, error) {

	query := `
		SELECT
			c.id,
			c.name,
			u.first_name || ' ' || u.last_name AS teacher_name,
			COUNT(DISTINCT sc.student_id) AS students_count,
			COUNT(DISTINCT cs.subject_id) AS subjects_count
		FROM classrooms c
		LEFT JOIN users u ON c.teacher_id = u.id
		LEFT JOIN student_classrooms sc ON sc.classroom_id = c.id
		LEFT JOIN classroom_subjects cs ON cs.classroom_id = c.id
		GROUP BY c.id, u.first_name, u.last_name
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classrooms []models.ClassroomCard

	for rows.Next() {
		var c models.ClassroomCard

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.TeacherName,
			&c.StudentsCount,
			&c.SubjectsCount,
		)

		if err != nil {
			return nil, err
		}

		classrooms = append(classrooms, c)
	}

	return classrooms, nil
}

func (r *ClassroomRepo) AddStudentsToClassroom(ctx context.Context, classroomID string, studentIDs []string) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, studentID := range studentIDs {

		// verify student exists
		var exists bool
		err := tx.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = 'student')`,
			studentID,
		).Scan(&exists)

		if err != nil {
			return err
		}

		if !exists {
			return errors.New("student not found: " + studentID)
		}

		_, err = tx.Exec(ctx,
			`INSERT INTO student_classrooms (classroom_id, student_id)
			 VALUES ($1, $2)
			 ON CONFLICT DO NOTHING`,
			classroomID,
			studentID,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *ClassroomRepo) GetClassroomByID(ctx context.Context, classroomID string) (*models.ClassroomDetail, error) {

	var detail models.ClassroomDetail

	// get classroom base info
	err := r.db.QueryRow(ctx, `
		SELECT
			c.id,
			c.name,
			u.first_name || ' ' || u.last_name AS teacher_name,
			c.academic_year
		FROM classrooms c
		LEFT JOIN users u ON c.teacher_id = u.id
		WHERE c.id = $1
	`, classroomID).Scan(
		&detail.ID,
		&detail.Name,
		&detail.TeacherName,
		&detail.AcademicYear,
	)

	if err != nil {
		return nil, errors.New("classroom not found")
	}

	// get subjects
	subjectRows, err := r.db.Query(ctx, `
		SELECT s.name
		FROM classroom_subjects cs
		JOIN subjects s ON cs.subject_id = s.id
		WHERE cs.classroom_id = $1
		ORDER BY s.name
	`, classroomID)

	if err != nil {
		return nil, err
	}
	defer subjectRows.Close()

	detail.Subjects = []string{}

	for subjectRows.Next() {
		var name string
		if err := subjectRows.Scan(&name); err != nil {
			return nil, err
		}
		detail.Subjects = append(detail.Subjects, name)
	}

	// get students
	studentRows, err := r.db.Query(ctx, `
		SELECT
			u.id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone
		FROM student_classrooms sc
		JOIN users u ON sc.student_id = u.id
		WHERE sc.classroom_id = $1
		ORDER BY u.first_name
	`, classroomID)

	if err != nil {
		return nil, err
	}
	defer studentRows.Close()

	detail.Students = []models.StudentInClass{}

	for studentRows.Next() {
		var s models.StudentInClass
		if err := studentRows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.Email,
			&s.Phone,
		); err != nil {
			return nil, err
		}
		detail.Students = append(detail.Students, s)
	}

	return &detail, nil
}

func (r *ClassroomRepo) RemoveStudentFromClassroom(ctx context.Context, classroomID string, studentID string) error {

	result, err := r.db.Exec(ctx,
		`DELETE FROM student_classrooms WHERE classroom_id = $1 AND student_id = $2`,
		classroomID,
		studentID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("student not found in this classroom")
	}

	return nil
}

func (r *ClassroomRepo) DeleteClassroom(ctx context.Context, classroomID string) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// delete from junction tables first
	_, err = tx.Exec(ctx,
		`DELETE FROM student_classrooms WHERE classroom_id = $1`, classroomID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`DELETE FROM classroom_subjects WHERE classroom_id = $1`, classroomID)
	if err != nil {
		return err
	}

	result, err := tx.Exec(ctx,
		`DELETE FROM classrooms WHERE id = $1`, classroomID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("classroom not found")
	}

	return tx.Commit(ctx)
}
