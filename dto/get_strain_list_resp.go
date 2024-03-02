package dto

type GetStrainResp struct {
	PageResp
	StrainList []Strain `json:"strain_list,omitempty"`
}
