package dto

type StrainListReq struct {
	PageReq
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}
