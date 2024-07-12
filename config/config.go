package config

import (
	"encoding/json"
	"github.com/linpanic/biology-server/utils"
)

type Config struct {
	Port    int64  `json:"port,omitempty"`     //web运行端口
	JWTKey  string `json:"jwt_key,omitempty"`  //JWT的key
	JWTTime int64  `json:"jwt_time,omitempty"` //jwt有效时间，单位秒
}

func LoadConfig(cfgPath string) *Config {
	cfgBytes := utils.ReadFile(cfgPath)
	c := new(Config)
	if len(cfgBytes) == 0 {
		return c
	}

	err := json.Unmarshal(cfgBytes, c)
	if err != nil {
		panic(err)
	}
	return c
}
