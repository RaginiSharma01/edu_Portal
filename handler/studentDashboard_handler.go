package handler

import (
	"log"
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type StudentDashboardHandler struct {
	service *service.StudentDashboardService
}

func NewStudentDashboardHandler(s *service.StudentDashboardService) *StudentDashboardHandler {
	return &StudentDashboardHandler{service: s}
}

func (h *StudentDashboardHandler) GetStudentDashboard(c fiber.Ctx) error {

	var req models.StudentDashboardRequest

	if err := c.Bind().Body(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	if req.StudentID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "studentId required",
		})
	}
	data, err := h.service.GetDashboard(c.Context(), req.StudentID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}
