package services

import (
	"context"

	"github.com/gui-marc/concurra/repository"
	"github.com/gui-marc/concurra/scheduler/models"
)

type EventService interface {
	GetEvents(ctx context.Context) ([]models.EventResponse, error)
	CreateEvent(ctx context.Context, name, startTime, endTime string, concurrencyTarget float32) (models.EventResponse, error)
	GetEvent(ctx context.Context, id string) (models.EventResponse, error)
	UpdateEvent(ctx context.Context, id, name, startTime, endTime string, concurrencyTarget float32) (models.EventResponse, error)
	DeleteEvent(ctx context.Context, id string) error
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
	events, err := s.eventRepository.GetEvents(ctx)
	if err != nil {
		return nil, err
	}

	var eventResponses []models.EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, models.NewEventResponse(
			event.ID,
			event.Name,
			event.StartTime,
			event.EndTime,
			event.ConcurrencyTarget,
		))
	}

	return eventResponses, nil
}

func (s *eventService) CreateEvent(ctx context.Context, name, startTime, endTime string, concurrencyTarget float32) (models.EventResponse, error) {
	id, err := s.eventRepository.CreateEvent(ctx, &repository.Event{
		Name:              name,
		StartTime:         startTime,
		EndTime:           endTime,
		ConcurrencyTarget: concurrencyTarget,
	})
	if err != nil {
		return models.EventResponse{}, err
	}

	return models.NewEventResponse(id, name, startTime, endTime, concurrencyTarget), nil
}

func (s *eventService) GetEvent(ctx context.Context, id string) (models.EventResponse, error) {
	event, err := s.eventRepository.GetEvent(ctx, id)
	if err != nil {
		return models.EventResponse{}, err
	}

	return models.NewEventResponse(event.ID, event.Name, event.StartTime, event.EndTime, event.ConcurrencyTarget), nil
}

func (s *eventService) UpdateEvent(ctx context.Context, id, name, startTime, endTime string, concurrencyTarget float32) (models.EventResponse, error) {
	err := s.eventRepository.UpdateEvent(ctx, id, &repository.Event{
		Name:              name,
		StartTime:         startTime,
		EndTime:           endTime,
		ConcurrencyTarget: concurrencyTarget,
	})
	if err != nil {
		return models.EventResponse{}, err
	}

	return models.NewEventResponse(id, name, startTime, endTime, concurrencyTarget), nil
}

func (s *eventService) DeleteEvent(ctx context.Context, id string) error {
	return s.eventRepository.DeleteEvent(ctx, id)
}
