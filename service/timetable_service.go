package service

import (
	"context"
	"smp/models"
	"smp/repository"
)

type TimetableService struct {
	Repo *repository.TimetableRepository
}

func NewTimetableService(repo *repository.TimetableRepository) *TimetableService {
	return &TimetableService{Repo: repo}
}

func (s *TimetableService) CreateTimetable(
	ctx context.Context,
	t models.CreateTimetable,
) error {

	return s.Repo.CreateTimeTable(ctx, t)
}
