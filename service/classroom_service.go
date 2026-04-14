package service

import (
	"context"
	"errors"
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

	if req.Name == "" {
		return "", errors.New("classroom name is required")
	}

	if req.TeacherID == "" {
		return "", errors.New("teacher ID is required")
	}

	if req.AcademicYear == "" {
		return "", errors.New("academic year is required")
	}

	if len(req.Subjects) == 0 {
		return "", errors.New("at least one subject is required")
	}

	return s.repo.CreateClassroom(ctx, req)
}

func (s *ClassroomService) GetClassrooms(ctx context.Context) ([]models.ClassroomCard, error) {
	return s.repo.GetClassrooms(ctx)
}

func (s *ClassroomService) AddStudentsToClassroom(ctx context.Context, classroomID string, studentIDs []string) error {

	if classroomID == "" {
		return errors.New("classroom ID is required")
	}

	if len(studentIDs) == 0 {
		return errors.New("at least one student ID is required")
	}

	return s.repo.AddStudentsToClassroom(ctx, classroomID, studentIDs)
}

func (s *ClassroomService) GetClassroomByID(ctx context.Context, classroomID string) (*models.ClassroomDetail, error) {

	if classroomID == "" {
		return nil, errors.New("classroom ID is required")
	}

	return s.repo.GetClassroomByID(ctx, classroomID)
}

func (s *ClassroomService) RemoveStudentFromClassroom(ctx context.Context, classroomID string, studentID string) error {

	if classroomID == "" || studentID == "" {
		return errors.New("classroom ID and student ID are required")
	}

	return s.repo.RemoveStudentFromClassroom(ctx, classroomID, studentID)
}

func (s *ClassroomService) DeleteClassroom(ctx context.Context, classroomID string) error {

	if classroomID == "" {
		return errors.New("classroom ID is required")
	}

	return s.repo.DeleteClassroom(ctx, classroomID)
}
