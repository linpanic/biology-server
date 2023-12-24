package model

// 品系名额外信息，一个品系可能对应多个额外信息
type StrainExtra struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	StrainId   int64  `json:"strain_id,omitempty"`   //品系名id
	ExtraKey   string `json:"extra_key,omitempty"`   //品系额外信息key
	ExtraValue string `json:"extra_value,omitempty"` //品系额外信息value
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (c *StrainExtra) TableName() string {
	return "strain_extra"
}
