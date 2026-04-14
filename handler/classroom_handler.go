package handler

import (
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

	id, err := h.service.CreateClassroom(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "classroom created successfully",
		"id":      id,
	})
}

func (h *ClassroomHandler) GetClassrooms(c fiber.Ctx) error {

	classrooms, err := h.service.GetClassrooms(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch classrooms",
		})
	}

	return c.JSON(classrooms)
}

func (h *ClassroomHandler) GetClassroomByID(c fiber.Ctx) error {

	classroomID := c.Params("id")

	detail, err := h.service.GetClassroomByID(c.Context(), classroomID)
	if err != nil {
		status := 500
		if err.Error() == "classroom not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(detail)
}

func (h *ClassroomHandler) AddStudentsToClassroom(c fiber.Ctx) error {

	classroomID := c.Params("id")

	var req models.AddStudentsToClassroom

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.service.AddStudentsToClassroom(c.Context(), classroomID, req.StudentIDs)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "students added to classroom successfully",
	})
}

func (h *ClassroomHandler) RemoveStudentFromClassroom(c fiber.Ctx) error {

	classroomID := c.Params("id")
	studentID := c.Params("studentId")

	err := h.service.RemoveStudentFromClassroom(c.Context(), classroomID, studentID)
	if err != nil {
		status := 500
		if err.Error() == "student not found in this classroom" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "student removed from classroom successfully",
	})
}

func (h *ClassroomHandler) DeleteClassroom(c fiber.Ctx) error {

	classroomID := c.Params("id")

	err := h.service.DeleteClassroom(c.Context(), classroomID)
	if err != nil {
		status := 500
		if err.Error() == "classroom not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "classroom deleted successfully",
	})
}
