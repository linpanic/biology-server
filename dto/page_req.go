package dto

type PageReq struct {
	PageNo   int `json:"page_no,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}
