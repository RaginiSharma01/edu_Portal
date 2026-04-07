package service

import (
	"context"
	"smp/models"
	"smp/repository"
)

type EventService struct {
	eventRepo *repository.EventRepo
}

func NewEventService(repo *repository.EventRepo) *EventService {
	return &EventService{
		eventRepo: repo,
	}
}

func (s *EventService) CreateEvent(ctx context.Context, req models.CreateEvent) (string, error) {
	return s.eventRepo.CreateEvent(ctx, req)

}

func (s *EventService) GetEvents(ctx context.Context) ([]models.Event, error) {
	return s.eventRepo.GetEvents(ctx)
}
