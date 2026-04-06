package models

type VerifyOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
