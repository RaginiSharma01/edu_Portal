package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOTPEmail(toEmail, otp string) error {

	from := mail.NewEmail("Student Portal", "raginisharma.r07@gmail.com")
	subject := "Email Verification OTP"

	// Load HTML template
	tmpl, err := template.ParseFiles("template/email_template.html")
	if err != nil {
		return err
	}

	data := struct {
		Email string
		OTP   string
	}{
		Email: toEmail,
		OTP:   otp,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	to := mail.NewEmail("", toEmail)

	message := mail.NewSingleEmail(
		from,
		subject,
		to,
		"Your OTP is: "+otp,
		body.String(),
	)

	client := sendgrid.NewSendClient(os.Getenv("API_KEYS"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	fmt.Println("SendGrid Status:", response.StatusCode)

	return nil
}
