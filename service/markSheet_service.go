package service

import (
	"context"
	"smp/models"
	"smp/repository"
)

type MarksService struct {
	repo *repository.MarksRepository
}

func NewMarksService(repo *repository.MarksRepository) *MarksService {
	return &MarksService{
		repo: repo,
	}
}

func (s *MarksService) CreateMarks(
	ctx context.Context,
	m models.CreateMarks,
) error {

	return s.repo.CreateMarks(ctx, m)
}

func (s *MarksService) GetMarks(
	ctx context.Context,
	term string,
) ([]models.StudentMarks, error) {

	return s.repo.GetMarks(ctx, term)
}
