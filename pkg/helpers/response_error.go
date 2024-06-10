package helpers

import (
	"log"
	"net/http"
)

type ResponseError struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Err     error          `json:"-"`
	Message MultiLanguages `json:"message"`
}

func NewError(code TypeError, err error) *ResponseError {

	var respErr MultiLanguages
	if errValue, isMatch := err.(*MultiLanguages); isMatch {
		if errValue != nil {
			respErr = *errValue
		} else {
			respErr = MultiLanguages{
				ID: err.Error(),
				EN: err.Error(),
			}
		}
	} else {
		respErr = MultiLanguages{
			ID: err.Error(),
			EN: err.Error(),
		}
	}
	return &ResponseError{
		Status:  StatusMapping(int(code)),
		Code:    int(code),
		Err:     err,
		Message: respErr,
	}
}

func (e *ResponseError) Error() string {
	return e.Err.Error()
}

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

func PanicRecover(opName string) {
	if r := recover(); r != nil {
		log.Printf("%v panic recover: %v \n", opName, r)
	}
}
