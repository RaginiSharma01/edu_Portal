package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) OnboardTeacher(c fiber.Ctx) error {

	var user models.TeacherOnboarding

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	userID, err := h.Service.OnboardUsers(c.Context(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "OTP sent to email",
		"user_id": userID,
	})
}
func (h *UserHandler) VerifyOTP(c fiber.Ctx) error {

	var req models.VerifyOTPRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.Service.VerifyOTP(c.Context(), req.Email, req.OTP)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "email verified successfully",
	})
}

func (h *UserHandler) Login(c fiber.Ctx) error {

	var req models.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	token, err := h.Service.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
	})
}
func (h *UserHandler) OnboardStudent(c fiber.Ctx) error {

	var student models.StudentOnboarding

	if err := c.Bind().Body(&student); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	userID, err := h.Service.OnboardStudent(c.Context(), student)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "OTP sent to email",
		"user_id": userID,
	})
}
func (h *UserHandler) GetAllTeachers(c fiber.Ctx) error {

	teachers, err := h.Service.GetAllTeachers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(teachers)
}

func (h *UserHandler) GetAllStudents(c fiber.Ctx) error {

	students, err := h.Service.GetAllStudents(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(students)
}
func (h *UserHandler) DeleteStudent(c fiber.Ctx) error {

	studentID := c.Params("id")

	err := h.Service.DeleteStudent(c.Context(), studentID)
	if err != nil {
		status := 500
		if err.Error() == "student not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "student deleted successfully",
	})
}

func (h *UserHandler) UpdateStudent(c fiber.Ctx) error {

	studentID := c.Params("id")

	var data models.UpdateStudent

	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.Service.UpdateStudent(c.Context(), studentID, data)
	if err != nil {
		status := 500
		if err.Error() == "student not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "student updated successfully",
	})
}

// BlockStudent handles both block and unblock via ?action=block or ?action=unblock
func (h *UserHandler) BlockStudent(c fiber.Ctx) error {

	studentID := c.Params("id")
	action := c.Query("action") // "block" or "unblock"

	if action != "block" && action != "unblock" {
		return c.Status(400).JSON(fiber.Map{
			"error": "query param 'action' must be 'block' or 'unblock'",
		})
	}

	block := action == "block"

	err := h.Service.BlockStudent(c.Context(), studentID, block)
	if err != nil {
		status := 500
		if err.Error() == "student not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	msg := "student blocked successfully"
	if !block {
		msg = "student unblocked successfully"
	}

	return c.JSON(fiber.Map{
		"message": msg,
	})
}
func (h *UserHandler) DeleteTeacher(c fiber.Ctx) error {

	teacherID := c.Params("id")

	err := h.Service.DeleteTeacher(c.Context(), teacherID)
	if err != nil {
		status := 500
		if err.Error() == "teacher not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "teacher deleted successfully",
	})
}

func (h *UserHandler) UpdateTeacher(c fiber.Ctx) error {

	teacherID := c.Params("id")

	var data models.UpdateTeacher

	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.Service.UpdateTeacher(c.Context(), teacherID, data)
	if err != nil {
		status := 500
		if err.Error() == "teacher not found" {
			status = 404
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "teacher updated successfully",
	})
}
