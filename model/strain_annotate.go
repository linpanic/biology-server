package model

// 品系名注解  一个品系可能对应多个注解
type StrainAnnotate struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	StrainId   int64  `json:"strain_id,omitempty"` //品系名id
	Annotate   string `json:"annotate,omitempty"`  //品系注解
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (c *StrainAnnotate) TableName() string {
	return "strain_annotate"
}
