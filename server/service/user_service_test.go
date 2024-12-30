package service

import (
	"bytes"
	"encoding/json"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	RegisterApi = "http://127.0.0.1:10080/register"
	LoginApi    = "http://127.0.0.1:10080/login"
	ListApi     = "http://127.0.0.1:10080/biology/strain_list"
	SearchApi   = "http://127.0.0.1:10080/biology/allele_search"
	DelApi      = "http://127.0.0.1:10080/biology/strain_delete"
)

func TestRegister(t *testing.T) {
	var req dto.UserRegisterReq
	req.Username = "test"
	req.Password = "abc123"
	req.Time = time.Now().Unix()

	buf := bytes.Buffer{}
	buf.WriteString(req.Username)
	timeStr := strconv.FormatInt(req.Time, 10)
	buf.WriteString(timeStr)
	buf.WriteString(req.Password)
	md5 := utils.MD5(buf.Bytes())
	md5 = strings.ToUpper(md5)
	req.Sign = md5
	marshal, _ := json.Marshal(req)
	log.Infof("请求api:%s,json为:%s", RegisterApi, string(marshal))
	resp := utils.HttpPostJson(RegisterApi, req)
	log.Info(string(resp.Data))
}

func TestLogin(t *testing.T) {
	var req dto.UserLoginReq
	req.Username = "test"
	req.Password = strings.ToUpper(utils.MD5([]byte(utils.MD5([]byte("abc123")))))
	req.Time = time.Now().Unix()
	buf := bytes.Buffer{}
	buf.WriteString(req.Username)
	timeStr := strconv.FormatInt(req.Time, 10)
	buf.WriteString(timeStr)
	buf.WriteString(req.Password)
	md5 := utils.MD5(buf.Bytes())
	md5 = strings.ToUpper(md5)
	req.Sign = md5
	marshal, _ := json.Marshal(req)
	log.Infof("请求api:%s,json为:%s", LoginApi, string(marshal))
	resp := utils.HttpPostJson(LoginApi, req)
	log.Info(string(resp.Data))
}

func TestGetList(t *testing.T) {
	h := make(map[string]string)
	h["X-Token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE2MjkyMjIsImlhdCI6MTcyMTAyNDQyMiwidXNlciI6Mn0.LK1ZjlD6STvbMGnk368NAa-h-bzxJgYIgTC1RJ2O538"
	resp := utils.HttpPostJsonAndHeader(ListApi, h, dto.StrainListReq{
		PageReq: dto.PageReq{
			PageNo:   19,
			PageSize: 5,
		},
		Key: "小米",
		//Field:   "number",
		//Order:   "desc",
	})
	log.Info(string(resp.Data))
}

func TestSearch(t *testing.T) {
	h := make(map[string]string)
	h["X-Token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE2MjkyMjIsImlhdCI6MTcyMTAyNDQyMiwidXNlciI6Mn0.LK1ZjlD6STvbMGnk368NAa-h-bzxJgYIgTC1RJ2O538"
	resp := utils.HttpPostJsonAndHeader(SearchApi, h, dto.AlleleSearchReq{
		Name: "拉",
	})
	log.Info(string(resp.Data))
}

func TestDel(t *testing.T) {
	h := make(map[string]string)
	h["X-Token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE2MjkyMjIsImlhdCI6MTcyMTAyNDQyMiwidXNlciI6Mn0.LK1ZjlD6STvbMGnk368NAa-h-bzxJgYIgTC1RJ2O538"
	resp := utils.HttpPostJsonAndHeader(DelApi, h, dto.StrainDelReq{
		StrainId: 1,
	})
	log.Info(string(resp.Data))
}
