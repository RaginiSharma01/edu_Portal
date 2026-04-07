package repository

import (
	"context"
	"fmt"
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

	fmt.Println(req)

	
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

	
	for _, subjectName := range req.Subjects {

		var subjectID string

		// Get subject id
		err := r.db.QueryRow(ctx,
			`SELECT id FROM subjects WHERE LOWER(name)=LOWER($1)`,
			subjectName,
		).Scan(&subjectID)

		if err != nil {
			return "", err
		}

		// Insert into classroom_subjects
		_, err = r.db.Exec(ctx,
			`INSERT INTO classroom_subjects (classroom_id, subject_id)
			 VALUES ($1, $2)`,
			classroomID,
			subjectID,
		)

		if err != nil {
			return "", err
		}
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
	GROUP BY c.id, u.first_name, u.last_name;
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
