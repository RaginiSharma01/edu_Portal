package service

import (
	"context"
	"smp/models"
	"smp/repository"
)

type TeacherDashboardService struct {
	repo *repository.TeacherDashboardRepository
}

func NewTeacherDashboardService(r *repository.TeacherDashboardRepository) *TeacherDashboardService {
	return &TeacherDashboardService{
		repo: r,
	}
}

func (s *TeacherDashboardService) GetDashboard(ctx context.Context, teacherID string) (*models.TeacherDashboard, error) {
	return s.repo.GetDashboard(ctx, teacherID)
}
