package dto

import "github.com/linpanic/biology-server/model"

type UserLoginResp struct {
	Token    string        `json:"token,omitempty"`
	UserName string        `json:"user_name,omitempty"`
	Menus    []*model.Menu `json:"menus,omitempty"`
}
