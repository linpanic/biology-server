package model

// 基因额外信息
type AlleleNameExtra struct {
	Id           int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	AlleleNameId int64  `json:"-"`                     //基因对应的ID
	ExtraKey     string `json:"extra_key,omitempty"`   //基因额外信息key
	ExtraValue   string `json:"extra_value,omitempty"` //基因额外信息value
	CreatorId    int64  `json:"creator_id,omitempty"`
	CreateTime   int64  `json:"create_time,omitempty"`
	UpdateTime   int64  `json:"update_time,omitempty"`
}

func (a *AlleleNameExtra) TableName() string {
	return "allele_name_extra"
}
