package handler

import (
	"fmt"
	"log"
	"smp/models"
	"smp/service"
	"smp/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler struct {
	Service *service.UserService
	DB      *pgxpool.Pool
}

func NewUserHandler(service *service.UserService, db *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		Service: service,
		DB:      db,
	}
}

func (h *UserHandler) OnboardTeacher(c fiber.Ctx) error {

	var user models.TeacherOnboarding

	if err := c.Bind().Body(&user); err != nil {
		log.Println("body parse err", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	userID, err := h.Service.OnboardUsers(c.Context(), user)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{

			"error": err.Error(),
		})
	}

	go utils.LogActivity(h.DB,
		fmt.Sprintf("New Teacher enrolled: %s", user.FirstName),
		"teacher",
	)

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

	// ADD THIS - log raw body to see what's arriving
	log.Println("Raw body:", string(c.Body()))

	var req models.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		log.Println("body parser error:", err) // this will show exact reason
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	// ADD THIS - log parsed struct
	log.Printf("Parsed login request: email=%s\n", req.Email)

	token, err := h.Service.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		log.Println(err)
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
		log.Println("body parser error:", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	log.Printf("Parsed student payload: %+v\n", student)

	userID, err := h.Service.OnboardStudent(c.Context(), student)
	if err != nil {
		log.Println("onboarding student service error:", err)
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Println("Student onboarded UserID", userID)

	go utils.LogActivity(h.DB,
		fmt.Sprintf("New Student has enrolled: %s", student.FirstName),
		"student",
	)

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
func (h *UserHandler) ForgotPassword(c fiber.Ctx) error {

	var req models.ForgotPasswordRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.Service.ForgotPassword(c.Context(), req.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "if this email exists, an OTP has been sent",
	})
}

func (h *UserHandler) ResetPassword(c fiber.Ctx) error {

	var req models.ResetPasswordRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.Service.ResetPassword(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "password reset successfully",
	})
}
