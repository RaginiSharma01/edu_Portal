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

	token, err := h.Service.VerifyOTP(c.Context(), req.Email, req.OTP)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "OTP verified",
		"token":   token,
	})
}
