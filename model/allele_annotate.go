package model

// 基因注释
type AlleleAnnotate struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	AlleleId   int64  `json:"-"`                  //对应的ID
	Annotate   string `json:"annotate,omitempty"` //注解
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (a *AlleleAnnotate) TableName() string {
	return "allele_annotate"
}
