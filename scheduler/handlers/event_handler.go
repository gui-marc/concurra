package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gui-marc/concurra/scheduler/models"
	"github.com/gui-marc/concurra/scheduler/services"
	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	GetEvents(c echo.Context) error
	CreateEvent(c echo.Context) error
	GetEventByID(c echo.Context) error
}

func NewEventHandler(eventService services.EventService) EventHandler {
	return &eventHandler{eventService: eventService}
}

type eventHandler struct {
	eventService services.EventService
}

func (h *eventHandler) GetEvents(c echo.Context) error {
	page := c.QueryParam("page")
	pageSize := c.QueryParam("pageSize")

	pageParams := models.PageParams{
		Page:     1,
		PageSize: 100,
	}

	if page != "" {
		if _, err := fmt.Sscanf(page, "%d", &pageParams.Page); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid page parameter"})
		}
	}

	if pageSize != "" {
		if _, err := fmt.Sscanf(pageSize, "%d", &pageParams.PageSize); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid pageSize parameter"})
		}
	}

	events, err := h.eventService.GetEvents(c.Request().Context(), pageParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, events)
}

func (h *eventHandler) CreateEvent(c echo.Context) error {
	type request struct {
		Name              string    `json:"name"`
		StartTime         time.Time `json:"startTime"`
		EndTime           time.Time `json:"endTime"`
		ConcurrencyTarget float32   `json:"concurrencyTarget"`
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

func (h *eventHandler) GetEventByID(c echo.Context) error {
	id := c.Param("eventId")

	event, err := h.eventService.GetEventByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, event)
}
