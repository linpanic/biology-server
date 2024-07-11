package model

//type StrainAllele struct {
//	StrainID         int64  `json:"strain_id"`
//	Number           string `json:"number"`
//	StrainName       string `json:"strain_name"`
//	StrainAnnotate   string `json:"strain_annotate"`
//	StrainExtraKey   string `json:"strain_extra_key"`
//	StrainExtraValue string `json:"strain_extra_value"`
//	ShortName        string `json:"short_name,omitempty"`
//	AlleleName       string `json:"allele_name"`
//	AlleleAnnotate   string `json:"allele_annotate"`
//	AlleleExtraKey   string `json:"allele_extra_key"`
//	AlleleExtraValue string `json:"allele_extra_value"`
//	GenomeName       string `json:"genome_name"`
//	Serial           string `json:"serial"`
//}

type StrainAllele struct {
	Id               int64  `json:"id,omitempty"`
	Number           string `json:"number"`
	StrainName       string `json:"strain_name"`
	AlleleId         string `json:"allele_id,omitempty"`
	AlleleName       string `json:"allele_name"`
	Genome           string `json:"genome,omitempty"`
	Serial           string `json:"serial"`
	AlleleExtraKey   string `json:"allele_extra_key"`
	AlleleExtraValue string `json:"allele_extra_value"`
	AAnnotate        string `json:"a_annotate,omitempty"`
	StrainAnnotate   string `json:"strain_annotate"`
	StrainExtraKey   string `json:"strain_extra_key"`
	StrainExtraValue string `json:"strain_extra_value"`
}
