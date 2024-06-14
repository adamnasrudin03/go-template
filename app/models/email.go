package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	SubjectEmailVerify   = "Verify Your Email Address"
	SubjectPasswordReset = "Reset Your Password"
)

// GenerateMessageEmailVerify generates the message for email verification
func GenerateMessageEmailVerify(userName, otp string, expired time.Duration) string {
	var sb strings.Builder
	sb.WriteString("Hi " + userName + ",\n\n")
	sb.WriteString("Thank you for using our service!\n")

	sb.WriteString("Verify your email to complete your registration.\n\n")
	sb.WriteString("Click here to verify: [Verification Link]\n")
	sb.WriteString("Your OTP code: " + otp + "\n")
	sb.WriteString(fmt.Sprintf("Code expires in: %v minutes\n", expired.Minutes()))
	sb.WriteString("\nIf you have trouble, contact support: [support email address]\n")
	sb.WriteString("Thank you for choosing our service!\n\n")
	sb.WriteString("Regards,\n\nThe [Company Name] Team")
	return sb.String()
}

func GenerateMessagePasswordReset(userName, otp string, expired time.Duration) string {
	var sb strings.Builder
	sb.WriteString("Hi " + userName + ",\n\n")
	sb.WriteString("Looks like you forgot your password for [App Name]. No worries, it happens!")
	sb.WriteString("\nTo set a new one, just:\n")
	sb.WriteString("1. Click here: [Password Reset Link]\n")
	sb.WriteString("2. Enter the code we sent to your email: " + otp + "\n")
	sb.WriteString(fmt.Sprintf("- Code expires in: %v minutes\n", expired.Minutes()))
	sb.WriteString("3. Pick a strong new password (something you haven't used before)\n")
	sb.WriteString("That's it! Your password will be reset, and you'll be back in action in no time.\n\n")
	sb.WriteString("If you didn't ask for a new password, you can just ignore this email.\n\n")
	sb.WriteString("See you soon,\n\n")
	sb.WriteString("The [Company Name] Team")
	return sb.String()
}
