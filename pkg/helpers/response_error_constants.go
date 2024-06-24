package helpers

type TypeError uint16

const (
	ErrForbidden    TypeError = iota + 10 // 10
	ErrUnauthorized                       // 11
	ErrDatabase                           // 12
	ErrConflict                           // 13
	ErrFromUseCase                        // 14
	ErrValidation                         // 15
	ErrNoFound                            // 16
	ErrUnknown                            // 17
)
