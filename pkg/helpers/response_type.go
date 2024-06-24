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
