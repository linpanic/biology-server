package dto

type Allele struct {
	Id       int64       `json:"id,omitempty"`
	Name     string      `json:"name"`               //基因名
	Genome   string      `json:"genome,omitempty"`   //基因修饰情况
	Serial   string      `json:"serial,omitempty"`   //对应第几条染色体
	Annotate []string    `json:"annotate,omitempty"` //基因修饰情况注解
	Extra    []ExtraInfo `json:"extra,omitempty"`    //基因修饰情况额外信息
}
