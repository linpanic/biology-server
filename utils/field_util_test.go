package utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestField(t *testing.T) {
	//var u dto.UserRegisterReq
	//u.Username = "1"
	//u.Password = ""
	//u.Time = 1
	//u.Sign = "x"
	//fmt.Println(FieldEmpty(&u)) //out:false
}

func TestJWT(t *testing.T) {
	jwt, err := GenJWT(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
	token, err := ParseToken(jwt)
	fmt.Println(token)
	fmt.Println(err)

	//resp := HttpGet("https://users.qzone.qq.com/fcg-bin/cgi_get_portrait.fcg?uins=694967189")
	//encode := GetStrEncode(resp.Data)
	//fmt.Println(resp.Data)
	//fmt.Println(encode)
	//fmt.Println(GBK2UTF8(resp.Data))
	//239 191 189 239 191 189 239 191 189 239 191 189

}

const (
	UNKNOW = -1
	GBK    = 1
	UTF8   = 2
)

func GetStrEncode(strBytes []byte) int {
	switch {
	case isGBK(strBytes):
		return GBK
	case isUtf8(strBytes):
		return UTF8
	}
	return UNKNOW
}

// GBK编码格式判断
func isGBK(data []byte) bool {
	length := len(data)
	var i = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0x7f {
				i = i + 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// UTF-8编码格式的判断
func preNUm(data byte) int {
	var mask byte = 0x80
	var num = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func GBK2UTF8(str []byte) string {
	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(str) //将gbk再转换为utf-8
	return string(utf8Data)
}
