package model

// 基因名 （一个品系名可能对应多个基因名，品系名可能为空）
type AlleleName struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	StrainId   int64  `json:"-"`                     //对应的品系名ID
	AlleleName string `json:"allele_name,omitempty"` //基因名 基因名是不重复的（但是可能有多个空）
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (a *AlleleName) TableName() string {
	return "allele_name"
}
