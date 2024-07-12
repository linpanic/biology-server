package dto

type AlleleUpdateReq struct {
	Id     int64    `json:"id,omitempty"`     //品系ID
	Allele []Allele `json:"allele,omitempty"` //基因信息
}
