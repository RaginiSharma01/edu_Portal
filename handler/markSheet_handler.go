package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type MarksHandler struct {
	marksService *service.MarksService
}

func NewMarksHandler(service *service.MarksService) *MarksHandler {
	return &MarksHandler{
		marksService: service,
	}
}

func (h *MarksHandler) CreateMarks(c fiber.Ctx) error {

	var req models.CreateMarks

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.marksService.CreateMarks(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "marks saved successfully",
	})
}

func (h *MarksHandler) GetMarks(c fiber.Ctx) error {

	term := c.Query("term")

	marks, err := h.marksService.GetMarks(c.Context(), term)
	if err != nil {
		return err
	}

	return c.JSON(marks)
}
