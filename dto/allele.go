package dto

type Allele struct {
	AlleleName         string      `json:"allele_name,omitempty"`          //基因名
	AlleleNameId       int64       `json:"allele_name_id,omitempty"`       //基因名ID
	AlleleNameAnnotate []string    `json:"allele_name_annotate,omitempty"` //基因修饰情况注解
	AlleleNameExtra    []ExtraInfo `json:"allele_name_extra,omitempty"`    //基因修饰情况额外信息
	GenomeId           int64       `json:"genome_id,omitempty"`            //基因修饰情况ID
	GenomeName         string      `json:"genome_name,omitempty"`          //基因修饰情况
	Serial             []Serial    `json:"serial,omitempty"`               //对应第几条染色体
}

type Serial struct {
	Id     int64  `json:"id,omitempty"`
	Serial string `json:"serial,omitempty"`
}
