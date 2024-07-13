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
)

func TestRegister(t *testing.T) {
	var req dto.UserRegisterReq
	req.Username = "buhuang2"
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
	req.Username = "buhuang2"
	req.Password = utils.MD5([]byte(utils.MD5([]byte("abc123"))))
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
