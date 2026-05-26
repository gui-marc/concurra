package handlers

import (
	"net/http"

	"github.com/gui-marc/concurra/scheduler/services"
	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	GetEvents(c echo.Context) error
	CreateEvent(c echo.Context) error
	GetEvent(c echo.Context) error
	UpdateEvent(c echo.Context) error
	DeleteEvent(c echo.Context) error
}

func NewEventHandler(eventService services.EventService) EventHandler {
	return &eventHandler{eventService: eventService}
}

type eventHandler struct {
	eventService services.EventService
}

func (h *eventHandler) GetEvents(c echo.Context) error {
	events, err := h.eventService.GetEvents(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, events)
}

func (h *eventHandler) CreateEvent(c echo.Context) error {
	type request struct {
		Name              string  `json:"name"`
		StartTime         string  `json:"startTime"`
		EndTime           string  `json:"endTime"`
		ConcurrencyTarget float32 `json:"concurrencyTarget"`
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid request body"})
	}

	event, err := h.eventService.CreateEvent(c.Request().Context(), req.Name, req.StartTime, req.EndTime, req.ConcurrencyTarget)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, event)
}

func (h *eventHandler) GetEvent(c echo.Context) error {
	id := c.Param("eventId")

	event, err := h.eventService.GetEvent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, event)
}

func (h *eventHandler) UpdateEvent(c echo.Context) error {
	id := c.Param("eventId")
	type request struct {
		Name              string  `json:"name"`
		StartTime         string  `json:"startTime"`
		EndTime           string  `json:"endTime"`
		ConcurrencyTarget float32 `json:"concurrencyTarget"`
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid request body"})
	}

	event, err := h.eventService.UpdateEvent(c.Request().Context(), id, req.Name, req.StartTime, req.EndTime, req.ConcurrencyTarget)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, event)
}

func (h *eventHandler) DeleteEvent(c echo.Context) error {
	id := c.Param("eventId")

	err := h.eventService.DeleteEvent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
