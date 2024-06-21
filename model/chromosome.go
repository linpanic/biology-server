package model

// 染色体信息 一个基因修饰情况可能有多个染色体信息
type Chromosome struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	GenomeId   int64  `json:"-"`                //基因修饰情况对应的ID
	Serial     string `json:"serial,omitempty"` //代表的是在第几条染色体上，用罗马数字标识
	CreatorId  int64  `json:"creator_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func (c *Chromosome) TableName() string {
	return "chromosome"
}
