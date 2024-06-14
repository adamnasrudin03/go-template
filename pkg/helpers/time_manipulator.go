package helpers

import (
	"fmt"
	"time"
)

// ParseUTC7 ...
func ParseUTC7(timeFormat string, value string) (time.Time, error) {
	timeUTC7, err := time.ParseInLocation(timeFormat, value, loc)
	if err != nil {
		return time.Now(), err
	}

	return timeUTC7, nil
}

// it will return 2009-11-10 00:00:00
func StartDate(t time.Time) time.Time {
	time, _ := ParseUTC7(FormatDate, t.Format(FormatDate))
	return time
}

// it will return ex 2009-11-10 23:59:59
func EndDate(t time.Time) time.Time {
	return StartDate(t).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

// it will return 2009-11-10 00:00:00
func StartDateString(t string) time.Time {
	time, _ := ParseUTC7(FormatDate, t)
	return time
}

// it will return ex 2009-11-10 23:59:59
func EndDateString(t string) time.Time {
	return StartDateString(t).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

func ConvertDurationToString(d time.Duration) string {
	res := ""
	if d.Hours() > 0 {
		res = res + fmt.Sprintf("%v hours ", d.Hours())
	}

	if d.Minutes() > 0 {
		res = res + fmt.Sprintf("%v minutes ", d.Minutes())
	}

	if d.Seconds() > 0 {
		res = res + fmt.Sprintf("%v seconds ", d.Seconds())
	}

	return res
}
