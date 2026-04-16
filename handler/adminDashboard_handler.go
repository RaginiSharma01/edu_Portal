package handler

import (
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type DashboardHandler struct {
	service *service.DashboardService
}

func NewDashboardHandler(service *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetAdminDashboard(c fiber.Ctx) error {

	data, err := h.service.GetAdminDashboard()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}
