package dto

type GetStrainReq struct {
	PageReq
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}
