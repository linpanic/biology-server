package model

// 基因修饰情况
// 如果有一个基因名，那么在不违规的情况下，一定有唯一一个确定的基因修饰情况
// 一个基因修饰情况，大部分情况下有一个基因名，但是绝不可能多个基因修饰情况对应一个基因名
type Genome struct {
	Id          int64  `json:"id,omitempty" gorm:"primary_key;not null;type:integer PRIMARY KEY AUTOINCREMENT"`
	AllelNameId int64  `json:"-"`                     //对应的ID
	GenomeName  string `json:"genome_name,omitempty"` //基因修饰情况  基因修饰情况有可能重复
	CreatorId   int64  `json:"creator_id,omitempty"`
	CreateTime  int64  `json:"create_time,omitempty"`
	UpdateTime  int64  `json:"update_time,omitempty"`
}

func (c *Genome) TableName() string {
	return "genome"
}
