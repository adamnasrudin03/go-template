package helpers

type ResponseDefault struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Meta struct
type Meta struct {
	Page         int `json:"page,omitempty"`
	Limit        int `json:"limit,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
}

type Pagination struct {
	Meta Meta
	Data interface{}
}

type MultiLanguages struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

func (e *MultiLanguages) Error() string {
	if e.EN != "" {
		return e.EN
	} else if e.ID != "" {
		return e.ID
	}
	return "something went wrong"
}

func NewResponseMultiLang(languages MultiLanguages) *MultiLanguages {
	return &languages
}
