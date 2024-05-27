package helpers

import (
	"encoding/json"
	"net/mail"
	"strings"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

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

// Round float generator
// example here: https://go.dev/play/p/j4B4v01Qolw
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
