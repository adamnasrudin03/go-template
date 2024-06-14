package models

import (
	"fmt"
	"time"
)

const (
	TimeDurationZero = time.Duration(0) * time.Second
	TimeDuration5Min = time.Duration(5) * time.Minute

	// key
	cacheUserDetail = "cache-user-detail"
	cacheOtp        = "cache-otp"
)

func GenerateKeyCacheUserDetail(userID uint64) string {
	return fmt.Sprintf("%s-%d", cacheUserDetail, userID)
}

func GenerateKeyCacheOtp(userID uint64, requestID string) string {
	return fmt.Sprintf("%s-user-%d-request-%s", cacheOtp, userID, requestID)
}
