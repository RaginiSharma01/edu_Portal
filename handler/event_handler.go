package handler

import (
	"fmt"
	"smp/models"
	"smp/service"
	"smp/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventHandler struct {
	eventService *service.EventService
	DB           *pgxpool.Pool
}

func NewEventHandler(s *service.EventService, db *pgxpool.Pool) *EventHandler {
	return &EventHandler{
		eventService: s,
		DB:           db,
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
	go utils.LogActivity(h.DB,
		fmt.Sprintf("Event created: %s", req.Title),
		"event",
	)

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
