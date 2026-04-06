package service

import (
	"context"
	"smp/models"
	"smp/repository"
)

type ClassroomService struct {
	repo *repository.ClassroomRepo
}

func NewClassroomService(repo *repository.ClassroomRepo) *ClassroomService {
	return &ClassroomService{repo: repo}
}

func (s *ClassroomService) CreateClassroom(ctx context.Context, req models.CreateClassroom) (string, error) {

	id, err := s.repo.CreateClassroom(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}
