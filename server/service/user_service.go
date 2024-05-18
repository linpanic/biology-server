package service

import (
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserService struct {
}

func (u UserService) Register(req dto.UserRegisterReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(err.Error())
	}
	count := dao.SelectUserCount(req.Username)
	if count != 0 {
		return dto.NewErrResult("用户已被注册")
	}
	tx := db.DbLink.Begin()
	err = dao.CreateUser(tx, req.Username, req.Password, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return dto.NewErrResult("注册失败")
	}
	tx.Commit()
	return dto.NewOKResult(nil)
}

func (u UserService) Login(req dto.UserLoginReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(err.Error())
	}
	user := dao.SelectOneUser(req.Username, req.Password)
	if user == nil {
		return dto.NewErrResult("用户名或者密码错误")
	}
	jwt, err := utils.GenJWT(user.Id)
	if err != nil {
		log.Error(err)
		return dto.NewErrResult("因未知原因登陆失败")
	}
	return dto.NewOKResult(dto.UserLoginResp{Token: jwt})
}

func (u UserService) ChangePassword(req dto.UserLoginReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(err.Error())
	}
	user := dao.SelectOneUser(req.Username, req.Password)
	if user == nil {
		return dto.NewErrResult("用户名或者密码错误")
	}
	jwt, err := utils.GenJWT(user.Id)
	if err != nil {
		log.Error(err)
		return dto.NewErrResult("因未知原因登陆失败")
	}
	return dto.NewOKResult(dto.UserLoginResp{Token: jwt})
}
