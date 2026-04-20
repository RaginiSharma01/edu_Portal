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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.marksService.CreateMarks(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "marks saved successfully",
	})
}

func (h *MarksHandler) GetMarks(c fiber.Ctx) error {

	term := c.Query("term")
	if term == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "term query param is required",
		})
	}

	marks, err := h.marksService.GetMarks(c.Context(), term)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(marks)
}

//  Download PDF for a single student by ID
func (h *MarksHandler) DownloadStudentPDF(c fiber.Ctx) error {

	studentID := c.Params("id")
	term := c.Query("term")

	if studentID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "student id param is required",
		})
	}

	if term == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "term query param is required",
		})
	}

	pdfBytes, err := h.marksService.GenerateStudentPDF(c.Context(), studentID, term)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=student_marksheet.pdf")

	return c.Send(pdfBytes)
}

// ✅ Download PDF for all students in a term
func (h *MarksHandler) DownloadMarksPDF(c fiber.Ctx) error {

	term := c.Query("term")

	if term == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "term query param is required",
		})
	}

	pdfBytes, err := h.marksService.GenerateMarksPDF(c.Context(), term)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=class_marksheet.pdf")

	return c.Send(pdfBytes)
}
