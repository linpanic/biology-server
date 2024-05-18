package dto

import (
	"bytes"
	"github.com/linpanic/biology-server/etype"
	"github.com/linpanic/biology-server/utils"
	"strconv"
	"strings"
	"time"
)

type UserChangePasswordReq struct {
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	Time        int64  `json:"time,omitempty"`
	Sign        string `json:"sign,omitempty"`
}

func (u *UserChangePasswordReq) Verify() error {
	isEmpty := utils.FieldEmpty(u)
	if isEmpty {
		return etype.FileNullError
	}
	if time.Now().Unix()-u.Time > 300 || u.Time-time.Now().Unix() > 300 {
		return etype.TimeError
	}
	buf := bytes.Buffer{}
	buf.WriteString(u.OldPassword)
	timeStr := strconv.FormatInt(u.Time, 10)
	buf.WriteString(timeStr)
	buf.WriteString(u.NewPassword)
	md5 := utils.MD5(buf.Bytes())
	md5 = strings.ToUpper(md5)
	if md5 != u.Sign {
		return etype.SignError
	}
	return nil
}
