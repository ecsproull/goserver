package services

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailRequest struct {
	To      string
	Subject string
	Text    string
	HTML    string
}

// SendEmail sends an email using SendGrid
func SendEmail(req EmailRequest) error {
	from := mail.NewEmail("", os.Getenv("SENDGRID_FROM_EMAIL"))
	to := mail.NewEmail("", req.To)

	// Use HTML if provided, otherwise fall back to text
	content := req.HTML
	if content == "" {
		content = req.Text
	}

	message := mail.NewSingleEmail(from, req.Subject, to, req.Text, content)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Printf("Email sent successfully: %d", response.StatusCode)
	return nil
}

// SendWelcomeEmail sends a welcome email to new users
func SendWelcomeEmail(userEmail, userName string) error {
	err := SendEmail(EmailRequest{
		To:      userEmail,
		Subject: "Welcome to Our Platform!",
		Text:    fmt.Sprintf("Hello %s, welcome to our platform!", userName),
		HTML: fmt.Sprintf(`
            <h1>Welcome %s!</h1>
            <p>Thank you for joining our platform.</p>
            <p>Best regards,<br>The Team</p>
						<p><b>If you need permissions to comment on blog posts or download files, reply to this email.</b></p>
        `, userName),
	})

	if err != nil {
		log.Printf("Failed to send welcome email: %v", err)
		return err
	}

	log.Printf("Welcome email sent to %s", userEmail)
	return nil
}

// SendPasswordResetEmail sends a password reset email
func SendPasswordResetEmail(userEmail, resetToken string) error {
	resetURL := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("FRONTEND_URL"), resetToken)

	err := SendEmail(EmailRequest{
		To:      userEmail,
		Subject: "Password Reset Request",
		Text:    fmt.Sprintf("Please click the following link to reset your password: %s", resetURL),
		HTML: fmt.Sprintf(`
            <h2>Password Reset Request</h2>
            <p>You requested a password reset. Click the link below to reset your password:</p>
            <a href="%s" style="background-color: #4CAF50; color: white; padding: 14px 20px; text-decoration: none; border-radius: 4px;">Reset Password</a>
            <p>If you didn't request this, please ignore this email.</p>
            <p>This link will expire in 1 hour.</p>
        `, resetURL),
	})

	if err != nil {
		log.Printf("Failed to send password reset email: %v", err)
		return err
	}

	log.Printf("Password reset email sent to %s", userEmail)
	return nil
}

// SendBlogNotification sends a notification email about new blog posts
func SendBlogNotification(userEmail, blogTitle, blogAuthor string) error {
	err := SendEmail(EmailRequest{
		To:      userEmail,
		Subject: fmt.Sprintf("New Blog Post: %s", blogTitle),
		Text:    fmt.Sprintf("A new blog post \"%s\" has been published by %s.", blogTitle, blogAuthor),
		HTML: fmt.Sprintf(`
            <h2>New Blog Post Published!</h2>
            <h3>%s</h3>
            <p>Author: %s</p>
            <p>Check out the latest blog post on our platform.</p>
        `, blogTitle, blogAuthor),
	})

	if err != nil {
		log.Printf("Failed to send blog notification: %v", err)
		return err
	}

	log.Printf("Blog notification sent to %s", userEmail)
	return nil
}

// SendVerificationEmail sends an email verification email
func SendVerificationEmail(userEmail, userName, verificationCode string) error {
	verificationURL := fmt.Sprintf("%s/verify-email?code=%s", os.Getenv("FRONTEND_URL"), verificationCode)

	err := SendEmail(EmailRequest{
		To:      userEmail,
		Subject: "Please verify your email address",
		Text:    fmt.Sprintf("Hello %s, please verify your email by clicking this link: %s", userName, verificationURL),
		HTML: fmt.Sprintf(`
            <h2>Welcome %s!</h2>
            <p>Thank you for signing up. Please verify your email address by clicking the button below:</p>
            <a href="%s" style="background-color: #4CAF50; color: white; padding: 14px 20px; text-decoration: none; border-radius: 4px; display: inline-block;">Verify Email</a>
            <p>Or copy and paste this link: %s</p>
            <p>This verification link will expire in 24 hours.</p>
        `, userName, verificationURL, verificationURL),
	})

	if err != nil {
		log.Printf("Failed to send verification email: %v", err)
		return err
	}

	log.Printf("Verification email sent to %s", userEmail)
	return nil
}
