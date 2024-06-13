package models

import "time"

const (
	TimeDurationZero = time.Duration(0) * time.Second
	TimeDuration5Min = time.Duration(5) * time.Minute

	// key
	CacheUserDetail = "cache-user-detail"
)
