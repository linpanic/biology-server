package dto

type Strain struct {
	Id             int64       `json:"id,omitempty"`
	StrainName     string      `json:"strain_name,omitempty"`
	Number         string      `json:"number,omitempty"`
	ShortName      []string    `json:"short_name,omitempty"`
	StrainAnnotate []string    `json:"strain_annotate,omitempty"`
	StrainExtra    []ExtraInfo `json:"strain_extra,omitempty"`
	Allele         []Allele    `json:"allele,omitempty"`
}

//type StrainShortName struct {
//	//Id        int64  `json:"id,omitempty"`
//	ShortName string `json:"short_name,omitempty"`
//}
//
//type StrainAnnotate struct {
//	//Id       int64  `json:"id,omitempty"`
//	Annotate string `json:"annotate,omitempty"`
//}
