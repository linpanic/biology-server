package model

type Menu struct {
	Id    uint64  `json:"id" gorm:"primaryKey;autoIncrement;not null"` //自增ID
	Pid   uint64  `json:"pid,omitempty"`                               //父级ID
	Name  string  `json:"name,omitempty"`                              //菜单名字
	Path  string  `json:"path,omitempty"`                              //菜单路径
	Child []*Menu `json:"child" gorm:"-"`
}

func (m *Menu) TableName() string {
	return "menu"
}
