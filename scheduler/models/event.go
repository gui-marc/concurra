package models

import (
	"time"

	"github.com/gui-marc/concurra/repository"
)

type EventResponse struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	StartTime         string  `json:"startTime"`
	EndTime           string  `json:"endTime"`
	ConcurrencyTarget float32 `json:"concurrencyTarget"`
}

func NewEventResponse(event *repository.Event) EventResponse {
	return EventResponse{
		ID:                event.ID,
		Name:              event.Name,
		StartTime:         event.StartTime.Format(time.RFC3339),
		EndTime:           event.EndTime.Format(time.RFC3339),
		ConcurrencyTarget: event.ConcurrencyTarget,
	}
}
