package handler

import (
	"log"
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
		log.Println("body parser err" , err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.salaryService.CreateSalary(c.Context(), req)
	if err != nil {
		log.Println(err)
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
func (h *SalaryHandler) UpdateSalary(c fiber.Ctx) error {

	teacherID := c.Params("teacherId")

	var data models.UpdateSalary

	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.salaryService.UpdateSalary(c.Context(), teacherID, data)
	if err != nil {
		status := 500
		if err.Error() == "salary record not found for this teacher" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "salary updated successfully",
	})
}
