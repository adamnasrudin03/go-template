package helpers

import (
	"net/http"
	"strings"
)

func StatusErrorMapping(code int) int {
	statusCode := 0
	switch code {
	case int(ErrForbidden):
		statusCode = http.StatusForbidden
	case int(ErrUnauthorized):
		statusCode = http.StatusUnauthorized
	case int(ErrDatabase):
		statusCode = http.StatusUnprocessableEntity
	case int(ErrFromUseCase):
		statusCode = http.StatusUnprocessableEntity
	case int(ErrConflict):
		statusCode = http.StatusConflict
	case int(ErrValidation):
		statusCode = http.StatusBadRequest
	case int(ErrNoFound):
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

// StatusMapping maps HTTP status code to a descriptive string.
// It returns the descriptive string that can be used as the 'status' field in ResponseDefault.
func StatusMapping(statusCode int) string {
	// Custom status codes that are not following standard HTTP status code text
	mappingsCustomStatus := map[int]string{
		http.StatusOK: "Success",
	}

	// Get the descriptive string for the custom status code
	status := mappingsCustomStatus[statusCode]

	// If the status code is not a custom status, get the standard descriptive string for the status code
	if status == "" {
		status = http.StatusText(statusCode)
	}

	// If the standard descriptive string is empty, the status code is not a standard HTTP status code
	// Map the status code to a HTTP status code for error responses and get the descriptive string for the error status code
	if status == "" {
		statusCode = StatusErrorMapping(statusCode)
		status = http.StatusText(statusCode)
	}

	status = strings.TrimSpace(status)
	return status
}

func StatusCodeMapping(statusCode int, v interface{}) int {
	if e, ok := v.(*ResponseError); ok {
		statusCode = StatusErrorMapping(e.Code)
	}
	return statusCode
}
