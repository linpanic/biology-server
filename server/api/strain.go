package api

import "github.com/linpanic/biology-server/server/service"

type StrainAPI struct {
	service.StrainService
}

var (
	StrainApi StrainAPI
)
