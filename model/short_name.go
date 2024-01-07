package model

// 简称  一个品系可能对应多个简称
type ShortName struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	StrainId   int64  `json:"strain_id,omitempty"`  //品系名id
	ShortName  string `json:"short_name,omitempty"` //简称
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (c *ShortName) TableName() string {
	return "short_name"
}
