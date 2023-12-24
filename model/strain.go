package model

// 品系名
type Strain struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	Number     string `json:"number,omitempty"`      //序列号，一般#开头
	StrainName string `json:"strain_name,omitempty"` //品系名(有人做出来之后没来得及命名 或者没有提交国际注册的话，可能是空)
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (s *Strain) TableName() string {
	return "strain"
}
