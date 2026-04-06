package handler

import (
	"context"
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type ClassroomHandler struct {
	service *service.ClassroomService
}

func NewClassroomHandler(service *service.ClassroomService) *ClassroomHandler {
	return &ClassroomHandler{service: service}
}

func (h *ClassroomHandler) CreateClassroom(c fiber.Ctx) error {

	var req models.CreateClassroom

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id, err := h.service.CreateClassroom(context.Background(), req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "classroom created",
		"id":      id,
	})
}
