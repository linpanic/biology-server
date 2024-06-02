package dto

type StrainListResp struct {
	PageResp
	StrainList []Strain `json:"strain_list,omitempty"`
}
