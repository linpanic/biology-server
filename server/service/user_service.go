package service

import (
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/model"
	"github.com/linpanic/biology-server/permission"
	"github.com/linpanic/biology-server/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type UserService struct {
}

func (u *UserService) Register(req dto.UserRegisterReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	count := dao.SelectUserCount(req.Username)
	if count != 0 {
		return dto.NewErrResult(cst.USER_EXIST, "用户已被注册")
	}
	tx := db.DbLink.Begin()
	req.Password = strings.ToUpper(utils.MD5([]byte(utils.MD5([]byte(req.Password)))))
	err = dao.CreateUser(tx, req.Username, req.Password, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, "注册失败")
	}
	tx.Commit()
	return dto.NewOKResult(nil)
}

func (u *UserService) Login(req dto.UserLoginReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	user := dao.SelectOneUser(req.Username, req.Password)
	if user == nil {
		return dto.NewErrResult(cst.PW_ERROR, "用户名或者密码错误")
	}
	jwt, err := utils.GenJWT(user.Id, user.UserName)
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.UNKNOW_ERROR, "因未知原因登陆失败")
	}
	menus := dao.GetMenus(db.DbLink)
	for i, v := range menus {
		if v.Path != "" {
			b, _ := permission.Ef.Enforce(user.UserName, v.Path, "read")
			if !b {
				menus[i] = nil
			}
		}
	}

	r := &model.Menu{
		Name: "root",
		Path: "/",
	}

	root := utils.GenNode(r, menus)
	utils.ClearNode(root)
	return dto.NewOKResult(dto.UserLoginResp{Token: jwt, UserName: req.Username, Menus: menus})
}

func (u *UserService) ChangePassword(req dto.UserLoginReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	user := dao.SelectOneUser(req.Username, req.Password)
	if user == nil {
		return dto.NewErrResult(cst.PW_ERROR, "用户名或者密码错误")
	}
	jwt, err := utils.GenJWT(user.Id, req.Username)
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.UNKNOW_ERROR, "因未知原因登陆失败")
	}
	return dto.NewOKResult(dto.UserLoginResp{Token: jwt})
}
