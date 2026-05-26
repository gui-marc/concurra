package repository

import (
	"context"
	"time"

	internal "github.com/gui-marc/concurra/repository/internal/generated/pgx"
)

type Event struct {
	ID                string
	Name              string
	StartTime         time.Time
	EndTime           time.Time
	ConcurrencyTarget float32
}

type CreateEventParams struct {
	Name              string
	StartTime         time.Time
	EndTime           time.Time
	ConcurrencyTarget float32
}

type EventRepository interface {
	GetEvents(ctx context.Context, limit int, offset int) ([]Event, error)
	GetEventByID(ctx context.Context, id string) (Event, error)
	CreateEvent(ctx context.Context, params CreateEventParams) (Event, error)
}

type pgxEventRepository struct {
	queries *internal.Queries
}

func NewPGXEventRepository(queries *internal.Queries) EventRepository {
	return &pgxEventRepository{
		queries: queries,
	}
}

func eventFromInternal(event internal.Event) Event {
	return Event{
		ID:                event.ID.String(),
		Name:              event.Name,
		StartTime:         event.StartTime.Time,
		EndTime:           event.EndTime.Time,
		ConcurrencyTarget: event.ConcurrencyTarget,
	}
}

func (r *pgxEventRepository) GetEvents(ctx context.Context, limit int, offset int) ([]Event, error) {
	events, err := r.queries.GetEvents(ctx, internal.GetEventsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	var result []Event
	for _, event := range events {
		result = append(result, eventFromInternal(event))
	}

	return result, nil
}

func (r *pgxEventRepository) GetEventByID(ctx context.Context, id string) (Event, error) {
	event, err := r.queries.GetEventByID(ctx, toPgUUID(id))
	if err != nil {
		return Event{}, err
	}

	return eventFromInternal(event), nil
}

func (r *pgxEventRepository) CreateEvent(ctx context.Context, params CreateEventParams) (Event, error) {
	event, err := r.queries.CreateEvent(ctx, internal.CreateEventParams{
		Name:              params.Name,
		StartTime:         toPgTimestamp(params.StartTime),
		EndTime:           toPgTimestamp(params.EndTime),
		ConcurrencyTarget: params.ConcurrencyTarget,
	})
	if err != nil {
		return Event{}, err
	}

	return eventFromInternal(event), nil
}
