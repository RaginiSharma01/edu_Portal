package service

import (
	"context"
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
