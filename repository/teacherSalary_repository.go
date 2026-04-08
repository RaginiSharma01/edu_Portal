package repository

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SalaryRepository struct {
	DB *pgxpool.Pool
}

func NewSalaryRepository(pool *pgxpool.Pool) *SalaryRepository {
	return &SalaryRepository{
		DB: pool,
	}
}

func (r *SalaryRepository) CreateSalary(ctx context.Context, salary models.CreateSalary) error {

	query := `
	INSERT INTO teacher_salaries
	(teacher_id, base_salary, allowance, effective_from)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.DB.Exec(
		ctx,
		query,

		salary.TeacherID,
		salary.BaseSalary,
		salary.Allowance,
		salary.EffectiveFrom,
	)

	return err
}

func (r *SalaryRepository) GetAllSalaries(ctx context.Context) ([]models.SalaryResponse, error) {

	query := `
	SELECT 
	u.first_name,
	u.last_name,
	s.base_salary,
	s.allowance,
	(s.base_salary + s.allowance) AS total,
	s.status
	FROM teacher_salaries s
	JOIN teachers t ON s.teacher_id = t.user_id
	JOIN users u ON u.id = t.user_id
	ORDER BY u.first_name
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salaries []models.SalaryResponse

	for rows.Next() {

		var firstName, lastName string
		var salary models.SalaryResponse

		err := rows.Scan(
			&firstName,
			&lastName,
			&salary.BaseSalary,
			&salary.Allowance,
			&salary.Total,
			&salary.Status,
		)

		if err != nil {
			return nil, err
		}

		salary.TeacherName = firstName + " " + lastName

		salaries = append(salaries, salary)
	}

	return salaries, nil
}
