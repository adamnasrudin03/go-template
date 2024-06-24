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
// The returned string can be used as the 'status' field in ResponseDefault.
func StatusMapping(statusCode int) string {
	mappingsCustomStatus := map[int]string{ // status custom not following standard http status code text
		http.StatusOK: "Success",
	}
	status := ""
	if statusCode >= 200 && statusCode < 300 {
		status = mappingsCustomStatus[statusCode]
		if status == "" {
			status = http.StatusText(statusCode)
		}
	}

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
