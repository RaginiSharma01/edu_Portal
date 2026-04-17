package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type TeacherDashboardHandler struct {
	service *service.TeacherDashboardService
}

// constructor
func NewTeacherDashboardHandler(s *service.TeacherDashboardService) *TeacherDashboardHandler {
	return &TeacherDashboardHandler{
		service: s,
	}
}

// handler function (using body as you wanted)
func (h *TeacherDashboardHandler) GetTeacherDashboard(c fiber.Ctx) error {
	var req models.DashboardRequest

	//
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// validation
	if req.TeacherID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "teacherId is required",
		})
	}

	// service call
	data, err := h.service.GetDashboard(c.Context(), req.TeacherID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}
