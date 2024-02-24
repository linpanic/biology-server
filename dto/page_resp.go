package dto

type PageResp struct {
	PageNo   int   `json:"page_no,omitempty"`
	PageSize int   `json:"page_size,omitempty"`
	Total    int64 `json:"total,omitempty"`
}
