package utils

import (
	"fmt"
	"github.com/linpanic/biology-server/dto"
	"testing"
)

func TestField(t *testing.T) {
	var u dto.UserRegisterReq
	u.Username = "1"
	u.Password = ""
	u.Time = 1
	u.Sign = "x"
	fmt.Println(FieldEmpty(&u)) //out:false
}
