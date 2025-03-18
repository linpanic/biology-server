package permission

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	log "github.com/sirupsen/logrus"
)

var (
	Ef *casbin.Enforcer
)

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
	a, _ := gormadapter.NewAdapter("sqlite3", "C:\\db\\biology.db", true) // Your driver and data source.
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatal(err)
	}
	Ef = e
	err = Ef.LoadPolicy()
	if err != nil {
		panic(err)
	}
}
