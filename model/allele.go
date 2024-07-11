package model

type Allele struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	StrainId   int64  `json:"-"`                //对应的品系名ID
	Name       string `json:"name,omitempty"`   //基因名
	Genome     string `json:"genome,omitempty"` //基因修饰情况
	Serial     string `json:"serial,omitempty"` //第几条染色体
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}
