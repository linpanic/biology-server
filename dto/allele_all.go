package dto

type AlleleAll struct {
	Id       int64       `json:"id,omitempty"`
	StrainId int64       `json:"strain_id,omitempty"`
	Name     string      `json:"name,omitempty"`     //基因名
	Genome   string      `json:"genome,omitempty"`   //基因修饰情况
	Serial   string      `json:"serial,omitempty"`   //对应第几条染色体
	Extra    []ExtraInfo `json:"extra,omitempty"`    //基因修饰情况额外信息
	Annotate []string    `json:"annotate,omitempty"` //基因修饰情况注解
}