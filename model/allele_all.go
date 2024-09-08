package model

// 基因全字段
type AlleleAll struct {
	AlleleId int64  `json:"allele_id,omitempty"`
	StrainId string `json:"strain_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Genome   string `json:"genome,omitempty"`
	Serial   string `json:"serial,omitempty"`
	ExtraKey string `json:"extra_key,omitempty"`
	ExtraVal string `json:"extra_val,omitempty"`
	Annotate string `json:"annotate,omitempty"`
}
