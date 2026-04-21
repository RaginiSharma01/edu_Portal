package service

import (
	"context"
	"smp/models"
	"smp/repository"
)

type StudentDashboardService struct {
	repo *repository.StudentDashboardRepository
}

func NewStudentDashboardService(r *repository.StudentDashboardRepository) *StudentDashboardService {
	return &StudentDashboardService{repo: r}
}

func (s *StudentDashboardService) GetDashboard(ctx context.Context, studentID string) (*models.StudentDashboard, error) {
	return s.repo.GetDashboard(ctx, studentID)
}