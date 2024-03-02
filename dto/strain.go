package dto

type Strain struct {
	StrainId       int64    `json:"strain_id,omitempty"`
	StrainName     string   `json:"strain_name,omitempty"`
	Number         string   `json:"number,omitempty"`
	ShortName      []string `json:"short_name,omitempty"`
	StrainAnnotate []string `json:"strain_annotate,omitempty"`
	StrainExtra    []string `json:"strain_extra,omitempty"`
	Allele         []Allele `json:"allele,omitempty"`
}
