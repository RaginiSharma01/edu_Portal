package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type SalaryHandler struct {
	salaryService *service.SalaryService
}

func NewSalaryHandler(service *service.SalaryService) *SalaryHandler {
	return &SalaryHandler{
		salaryService: service,
	}
}

func (h *SalaryHandler) CreateSalary(c fiber.Ctx) error {

	var req models.CreateSalary

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.salaryService.CreateSalary(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create salary",
		})
	}
	// err := h.salaryService.CreateSalary(c.Context(), req)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	return c.JSON(fiber.Map{
		"message": "salary created successfully",
	})
}
func (h *SalaryHandler) GetAllSalaries(c fiber.Ctx) error {

	salaries, err := h.salaryService.GetAllSalaries(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch salaries",
		})
	}

	return c.JSON(salaries)
}
