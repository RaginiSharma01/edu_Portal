package repository

import (
	"context"
	"errors"
	"smp/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		DB: pool,
	}
}
func (r *UserRepo) OnboardingUser(ctx context.Context, user models.TeacherOnboarding) (string, error) {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	var userID string

	Query := `
	INSERT INTO users
	(first_name, last_name, email, phone, age, date_of_birth, address, password, role)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id
	`

	err = tx.QueryRow(
		ctx,
		Query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Age,
		user.DateOfBirth,
		user.Address,
		user.Password,
		"teacher",
	).Scan(&userID)

	if err != nil {
		return "", err
	}

	teacherQuery := `
	INSERT INTO teachers
	(user_id, qualification, subjects_teaching)
	VALUES ($1,$2,$3)
	`

	_, err = tx.Exec(
		ctx,
		teacherQuery,
		userID,
		user.Qualification,
		user.SubjectsTeaching,
	)

	if err != nil {
		return "", err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	query := `
	SELECT id, email, password, role, email_verified
FROM users
WHERE email=$1
	`

	var user models.User

	err := r.DB.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsVerified,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) VerifyUser(ctx context.Context, email string) error {

	query := `
	UPDATE users
	SET email_verified = true
	WHERE email = $1
	`

	_, err := r.DB.Exec(ctx, query, email)
	return err
}

func (r *UserRepo) OnboardStudent(ctx context.Context, student models.StudentOnboarding) (string, error) {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	var userID string

	userQuery := `
	INSERT INTO users
	(first_name, last_name, email, phone, age, date_of_birth, address, password, role)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id
	`

	err = tx.QueryRow(
		ctx,
		userQuery,
		student.FirstName,
		student.LastName,
		student.Email,
		student.Phone,
		student.Age,
		student.DateOfBirth,
		student.Address,
		student.Password,
		"student",
	).Scan(&userID)

	if err != nil {
		return "", err
	}

	studentQuery := `
	INSERT INTO students
	(user_id, father_name, mother_name, guardian_name, occupation, height, weight)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err = tx.Exec(
		ctx,
		studentQuery,
		userID,
		student.FatherName,
		student.MotherName,
		student.GuardianName,
		student.Occupation,
		student.Height,
		student.Weight,
	)

	if err != nil {
		return "", err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// func to get student and get teachers
func (r *UserRepo) GetAllTeachers(ctx context.Context) (pgx.Rows, error) {

	query := `
	SELECT 
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.phone,
		u.age,
		u.date_of_birth,
		u.address,
		t.qualification,
		t.subjects_teaching
	FROM teachers t
	JOIN users u ON u.id = t.user_id
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *UserRepo) GetAllStudents(ctx context.Context) (pgx.Rows, error) {

	query := `
	SELECT 
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.phone,
		u.age,
		u.date_of_birth,
		u.address,
		s.father_name,
		s.mother_name,
		s.guardian_name,
		s.occupation,
		s.height,
		s.weight
	FROM students s
	JOIN users u ON u.id = s.user_id
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// DeleteStudent hard-deletes a student by user ID
func (r *UserRepo) DeleteStudent(ctx context.Context, studentID string) error {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// delete from students table first (FK constraint)
	_, err = tx.Exec(ctx, `DELETE FROM students WHERE user_id = $1`, studentID)
	if err != nil {
		return err
	}

	// delete from users table
	result, err := tx.Exec(ctx, `DELETE FROM users WHERE id = $1 AND role = 'student'`, studentID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("student not found")
	}

	return tx.Commit(ctx)
}

// UpdateStudent updates user + student profile details
func (r *UserRepo) UpdateStudent(ctx context.Context, studentID string, data models.UpdateStudent) error {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	userQuery := `
    UPDATE users SET
        first_name   = $1,
        last_name    = $2,
        phone        = $3,
        age          = $4,
        date_of_birth = $5,
        address      = $6
    WHERE id = $7 AND role = 'student'
    `

	result, err := tx.Exec(ctx, userQuery,
		data.FirstName,
		data.LastName,
		data.Phone,
		data.Age,
		data.DateOfBirth,
		data.Address,
		studentID,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("student not found")
	}

	studentQuery := `
    UPDATE students SET
        father_name   = $1,
        mother_name   = $2,
        guardian_name = $3,
        occupation    = $4,
        height        = $5,
        weight        = $6
    WHERE user_id = $7
    `

	_, err = tx.Exec(ctx, studentQuery,
		data.FatherName,
		data.MotherName,
		data.GuardianName,
		data.Occupation,
		data.Height,
		data.Weight,
		studentID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// BlockStudent toggles the is_blocked flag on a user
func (r *UserRepo) BlockStudent(ctx context.Context, studentID string, block bool) error {

	query := `
    UPDATE users
    SET is_blocked = $1
    WHERE id = $2 AND role = 'student'
    `

	result, err := r.DB.Exec(ctx, query, block, studentID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("student not found")
	}

	return nil
}

func (r *UserRepo) DeleteTeacher(ctx context.Context, teacherID string) error {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM teacher_salaries WHERE teacher_id = $1`, teacherID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `DELETE FROM teachers WHERE user_id = $1`, teacherID)
	if err != nil {
		return err
	}

	result, err := tx.Exec(ctx, `DELETE FROM users WHERE id = $1 AND role = 'teacher'`, teacherID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("teacher not found")
	}

	return tx.Commit(ctx)
}

func (r *UserRepo) UpdateTeacher(ctx context.Context, teacherID string, data models.UpdateTeacher) error {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	userQuery := `
    UPDATE users SET
        first_name    = $1,
        last_name     = $2,
        phone         = $3,
        age           = $4,
        date_of_birth = $5,
        address       = $6
    WHERE id = $7 AND role = 'teacher'
    `

	result, err := tx.Exec(ctx, userQuery,
		data.FirstName,
		data.LastName,
		data.Phone,
		data.Age,
		data.DateOfBirth,
		data.Address,
		teacherID,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("teacher not found")
	}

	teacherQuery := `
    UPDATE teachers SET
        qualification     = $1,
        subjects_teaching = $2
    WHERE user_id = $3
    `

	_, err = tx.Exec(ctx, teacherQuery,
		data.Qualification,
		data.SubjectsTeaching,
		teacherID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
