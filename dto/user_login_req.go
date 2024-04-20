package dto

import (
	"bytes"
	"github.com/linpanic/biology-server/etype"
	"github.com/linpanic/biology-server/utils"
	"strconv"
	"strings"
	"time"
)

type UserLoginReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Time     int64  `json:"time,omitempty"`
	Sign     string `json:"sign,omitempty"`
}

func (u *UserLoginReq) Verify() error {
	isEmpty := utils.FieldEmpty(u)
	if isEmpty {
		return etype.FileNullError
	}
	if time.Now().Unix()-u.Time > 300 || u.Time-time.Now().Unix() > 300 {
		return etype.TimeError
	}
	buf := bytes.Buffer{}
	buf.WriteString(u.Username)
	timeStr := strconv.FormatInt(u.Time, 10)
	buf.WriteString(timeStr)
	buf.WriteString(u.Password)
	md5 := utils.MD5(buf.Bytes())
	md5 = strings.ToUpper(md5)
	if md5 != u.Sign {
		return etype.SignError
	}
	return nil
}
