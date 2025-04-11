package model

type Pagination struct {
	Limit         int `json:"limit"`
	Offset        int `json:"offset"`
	RowCount      int `json:"rowCount"`
	TotalRowCount int `json:"totalRowCount"`
	Page          int `json:"page"`
	MaxPage       int `json:"maxPage"`
}

type GeneralResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
