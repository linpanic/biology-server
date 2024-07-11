package model

type StrainJoin struct {
	StrainID       int64  `json:"strain_id"`
	Number         string `json:"number"`
	StrainName     string `json:"strain_name"`
	StrainAnnotate string `json:"strain_annotate"`
	ExtraKey       string `json:"extra_key"`
	ExtraValue     string `json:"extra_value"`
	ShortName      string `json:"short_name"`
}
