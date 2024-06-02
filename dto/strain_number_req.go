package dto

import (
	"github.com/linpanic/biology-server/etype"
	"github.com/linpanic/biology-server/utils"
	"strconv"
	"strings"
	"time"
)

type StrainNumberReq struct {
	Time int64  `json:"time,omitempty"`
	Sign string `json:"sign,omitempty"`
}

func (s *StrainNumberReq) Verify() error {
	isEmpty := utils.FieldEmpty(s)
	if isEmpty {
		return etype.FileNullError
	}
	if time.Now().Unix()-s.Time > 300 || s.Time-time.Now().Unix() > 300 {
		return etype.TimeError
	}
	timeStr := strconv.FormatInt(s.Time, 10)
	md5 := utils.MD5([]byte(timeStr))
	md5 = strings.ToUpper(md5)
	if md5 != s.Sign {
		return etype.SignError
	}
	return nil
}
