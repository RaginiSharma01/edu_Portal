package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type TimetableHandler struct {
	timetableService *service.TimetableService
}

func NewTimetableHandler(service *service.TimetableService) *TimetableHandler {
	return &TimetableHandler{
		timetableService: service,
	}
}

func (h *TimetableHandler) CreateTimetable(c fiber.Ctx) error {

	var req models.CreateTimetable

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	err := h.timetableService.CreateTimetable(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "timetable created",
	})
}
func (h *TimetableHandler) GetTimetable(c fiber.Ctx) error {

	classID := c.Params("classId")

	timetable, err := h.timetableService.GetTimetable(c.Context(), classID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch timetable",
		})
	}

	return c.JSON(timetable)
}
