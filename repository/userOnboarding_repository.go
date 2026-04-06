package repository

import (
	"context"
	"smp/models"

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
