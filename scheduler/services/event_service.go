package services

import (
	"context"
	"time"

	"github.com/gui-marc/concurra/repository"
	"github.com/gui-marc/concurra/scheduler/models"
)

type EventService interface {
	GetEvents(ctx context.Context) ([]models.EventResponse, error)
	CreateEvent(ctx context.Context, name string, startTime, endTime time.Time, concurrencyTarget float32) (models.EventResponse, error)
	GetEventByID(ctx context.Context, id string) (models.EventResponse, error)
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{
		eventRepository: eventRepository,
	}
}

type eventService struct {
	eventRepository repository.EventRepository
}

func (s *eventService) GetEvents(ctx context.Context) ([]models.EventResponse, error) {
	events, err := s.eventRepository.GetEvents(ctx, 100, 0)
	if err != nil {
		return nil, err
	}

	var eventResponses []models.EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, models.NewEventResponse(&event))
	}

	return eventResponses, nil
}

func (s *eventService) CreateEvent(ctx context.Context, name string, startTime, endTime time.Time, concurrencyTarget float32) (models.EventResponse, error) {
	event, err := s.eventRepository.CreateEvent(ctx, repository.CreateEventParams{
		Name:              name,
		StartTime:         startTime,
		EndTime:           endTime,
		ConcurrencyTarget: concurrencyTarget,
	})
	if err != nil {
		return models.EventResponse{}, err
	}

	return models.NewEventResponse(&event), nil
}

func (s *eventService) GetEventByID(ctx context.Context, id string) (models.EventResponse, error) {
	event, err := s.eventRepository.GetEventByID(ctx, id)
	if err != nil {
		return models.EventResponse{}, err
	}

	return models.NewEventResponse(&event), nil
}
