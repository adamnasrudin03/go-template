package helpers

import "strings"

// unmarshal from response http
type ResponseErrorHttp struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Desc    MultiLanguages `json:"desc"`
	Message MultiLanguages `json:"message"`
}

func (m *ResponseErrorHttp) GetMessageID() string {
	message := strings.TrimSpace(m.Desc.ID)
	if message == "" {
		message = strings.TrimSpace(m.Message.ID)
	}
	return message
}

func (m *ResponseErrorHttp) GetMessageEN() string {
	message := strings.TrimSpace(m.Desc.EN)
	if message == "" {
		message = strings.TrimSpace(m.Message.EN)
	}
	return message
}
