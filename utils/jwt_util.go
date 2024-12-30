package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/linpanic/biology-server/caches"
	"time"
)

// 生成JWT
func GenJWT(userId int64, username string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Second * time.Duration(caches.JWTTime)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user"] = userId
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString(caches.JWTKey)
}

func secret() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return caches.JWTKey, nil
	}
}

// 解析token
func ParseToken(token string) (int64, error) {
	tokn, _ := jwt.Parse(token, secret())
	claim, ok := tokn.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("解析错误")
	}
	if !tokn.Valid {
		return 0, errors.New("令牌错误！")
	}
	exp := claim["exp"].(float64)
	if exp == 0 || int64(exp) < time.Now().Unix() {
		return 0, errors.New("用户登录过期")
	}
	return int64(claim["user"].(float64)), nil
}
