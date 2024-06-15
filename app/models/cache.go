package models

import (
	"fmt"
	"time"
)

const (
	TimeDurationZero = time.Duration(0) * time.Second
)

func GenerateKeyCacheUserDetail(userID uint64) string {
	return fmt.Sprintf("cache-user-detail-%d", userID)
}

func GenerateKeyCacheOtp(userID uint64, requestID string) string {
	return fmt.Sprintf("cache-otp-user-%d-request-%s", userID, requestID)
}

func GenerateKeyCacheForgotPassword(userID uint64, requestID string) string {
	return fmt.Sprintf("cache-forgot-password-user-%d-request-%s", userID, requestID)
}
