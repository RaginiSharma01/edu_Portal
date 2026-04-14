package service

import (
	"context"
	"errors"
	"smp/models"
	"smp/repository"
)

type SalaryService struct {
	salaryRepo *repository.SalaryRepository
}

func NewSalaryService(repo *repository.SalaryRepository) *SalaryService {
	return &SalaryService{
		salaryRepo: repo,
	}
}

func (s *SalaryService) CreateSalary(ctx context.Context, salary models.CreateSalary) error {
	return s.salaryRepo.CreateSalary(ctx, salary)
}

func (s *SalaryService) GetAllSalaries(ctx context.Context) ([]models.SalaryResponse, error) {
	return s.salaryRepo.GetAllSalaries(ctx)
}
func (s *SalaryService) UpdateSalary(ctx context.Context, teacherID string, data models.UpdateSalary) error {

	if teacherID == "" {
		return errors.New("teacher ID required")
	}

	if data.BaseSalary <= 0 {
		return errors.New("base salary must be greater than 0")
	}

	return s.salaryRepo.UpdateSalary(ctx, teacherID, data)
}
