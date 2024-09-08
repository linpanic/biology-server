package dto

type AlleleAllListResp struct {
	PageResp
	Allele []AlleleAll `json:"allele,omitempty"`
}
