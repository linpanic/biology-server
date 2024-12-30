package utils

import "github.com/linpanic/biology-server/model"

func ClearNode(m *model.Menu) {
	var newChild []*model.Menu
	for i, v := range m.Child {
		if len(v.Child) > 0 {
			ClearNode(v)
		}
		if v.Path == "" && len(v.Child) == 0 {
			continue
		} else {
			newChild = append(newChild, m.Child[i])
		}
	}
	m.Child = newChild
}

func GenNode(p *model.Menu, ls []*model.Menu) *model.Menu {
	for _, v := range ls {
		if v == nil {
			continue
		}
		if p.Id == v.Pid {
			node := GenNode(v, ls)
			p.Child = append(p.Child, node)
		}
	}
	return p
}
