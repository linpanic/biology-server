package permission

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	Ef *casbin.Enforcer
)

func init() {
	InitCasbin()
}

func InitCasbin() {
	m, err := model.NewModelFromString(`[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act, eft

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) 

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`)
	if err != nil {
		panic(err)
	}
	a, _ := gormadapter.NewAdapter("sqlite3", "C:\\db\\biology.db") // Your driver and data source.
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Ef = e
	err = Ef.LoadPolicy()
	if err != nil {
		panic(err)
	}

}
