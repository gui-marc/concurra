package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Event struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	StartTime         string  `json:"startTime"`
	EndTime           string  `json:"endTime"`
	ConcurrencyTarget float32 `json:"concurrencyTarget"`
}

type eventTable struct{}

func NewEventTable() Table[Event] {
	return &eventTable{}
}

func (t *eventTable) GetName() string {
	return "events"
}

func (t *eventTable) GetColumns() []string {
	return []string{"id", "name", "start_time", "end_time", "concurrency_target"}
}

func (t *eventTable) Read(record *Record) (*Event, error) {
	event := &Event{
		ID:                (*record)["id"].(string),
		Name:              (*record)["name"].(string),
		StartTime:         (*record)["start_time"].(string),
		EndTime:           (*record)["end_time"].(string),
		ConcurrencyTarget: (*record)["concurrency_target"].(float32),
	}
	return event, nil
}

func (t *eventTable) ToRecord(event *Event) *Record {
	return &Record{
		"id":                 event.ID,
		"name":               event.Name,
		"start_time":         event.StartTime,
		"end_time":           event.EndTime,
		"concurrency_target": event.ConcurrencyTarget,
	}
}

func newEventPGDatabase(conn *pgx.Conn) Database[Event] {
	return NewPostgresDatabase[Event](conn, NewEventTable())
}

type EventRepository interface {
	GetEvents(ctx context.Context) ([]*Event, error)
	CreateEvent(ctx context.Context, event *Event) (string, error)
	GetEvent(ctx context.Context, id string) (*Event, error)
	UpdateEvent(ctx context.Context, id string, event *Event) error
	DeleteEvent(ctx context.Context, id string) error
}

type eventRepository struct {
	db Database[Event]
}

func NewEventRepository(conn *pgx.Conn) EventRepository {
	return &eventRepository{
		db: newEventPGDatabase(conn),
	}
}

func (r *eventRepository) GetEvents(ctx context.Context) ([]*Event, error) {
	events, err := r.db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) CreateEvent(ctx context.Context, event *Event) (string, error) {
	event.ID = uuid.New().String()
	return r.db.Create(event)
}

func (r *eventRepository) GetEvent(ctx context.Context, id string) (*Event, error) {
	return r.db.Get(id)
}

func (r *eventRepository) UpdateEvent(ctx context.Context, id string, event *Event) error {
	return r.db.Update(id, event)
}

func (r *eventRepository) DeleteEvent(ctx context.Context, id string) error {
	return r.db.Delete(id)
}
