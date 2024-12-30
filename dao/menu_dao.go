package dao

import (
	"github.com/linpanic/biology-server/model"
	"gorm.io/gorm"
)

func GetMenus(dbLink *gorm.DB) []*model.Menu {
	var menus []*model.Menu
	err := dbLink.Model(&model.Menu{}).Find(&menus).Error
	if err != nil {
		panic(err)
	}
	return menus
}
