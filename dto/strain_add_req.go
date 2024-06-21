package dto

type StrainAddReq struct {
	StrainName     string      `json:"strain_name,omitempty"`
	Number         string      `json:"number,omitempty"`
	ShortName      []string    `json:"short_name,omitempty"`
	StrainAnnotate []string    `json:"strain_annotate,omitempty"`
	StrainExtra    []ExtraInfo `json:"strain_extra,omitempty"`
	Allele         []Allele    `json:"allele,omitempty"`
}
