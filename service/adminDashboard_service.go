package service

import "smp/repository"

type DashboardService struct {
	repo *repository.AdminDashboardRepository
}

func NewDashboardService(repo *repository.AdminDashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetAdminDashboard() (map[string]interface{}, error) {
	return s.repo.GetDashboard()
}