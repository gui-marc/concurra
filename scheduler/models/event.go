package models

type EventResponse struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	StartTime         string  `json:"startTime"`
	EndTime           string  `json:"endTime"`
	ConcurrencyTarget float32 `json:"concurrencyTarget"`
}

func NewEventResponse(id, name, startTime, endTime string, concurrencyTarget float32) EventResponse {
	return EventResponse{
		ID:                id,
		Name:              name,
		StartTime:         startTime,
		EndTime:           endTime,
		ConcurrencyTarget: concurrencyTarget,
	}
}
