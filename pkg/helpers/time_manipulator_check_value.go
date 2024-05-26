package helpers

import "time"

func CheckTimeIsZeroToPointer(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func CheckTimePointerValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func CheckTimeIsZeroToString(t time.Time, formatDate string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(formatDate)
}

func CheckTimePointerToString(t *time.Time, formatDate string) string {
	if t == nil {
		return ""
	}
	return t.Format(formatDate)
}
