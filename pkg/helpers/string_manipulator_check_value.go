package helpers

import (
	"strings"
)

// Return empty string if data is nil, otherwise return the value of data
func CheckStringValue(data *string) string {
	if data == nil {
		return ""
	}
	return *data
}

// Return a pointer to a string or nil if the input string is empty or only contains whitespace
func CheckStringValueToPointer(data string) *string {
	data = strings.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}
	return &data
}
