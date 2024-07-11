package dto

type Allele struct {
	AlleleName     string      `json:"allele_name,omitempty"`     //基因名
	GenomeName     string      `json:"genome_name,omitempty"`     //基因修饰情况
	Serial         string      `json:"serial,omitempty"`          //对应第几条染色体
	AlleleAnnotate []string    `json:"allele_annotate,omitempty"` //基因修饰情况注解
	Extra          []ExtraInfo `json:"extra,omitempty"`           //基因修饰情况额外信息
}
