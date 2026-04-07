package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type EventHandler struct {
	eventService *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{
		eventService: s,
	}
}

func (h *EventHandler) CreateEvent(c fiber.Ctx) error {

	var req models.CreateEvent

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	eventID, err := h.eventService.CreateEvent(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create event",
		})
	}

	return c.JSON(fiber.Map{
		"message":  "event created successfully",
		"event_id": eventID,
	})
}

func (h *EventHandler) GetEvents(c fiber.Ctx) error {

	events, err := h.eventService.GetEvents(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch events",
		})
	}

	return c.JSON(events)
}