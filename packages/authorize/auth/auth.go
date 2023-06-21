package auth

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2/model"
	"github.com/example-golang-projects/clean-architecture/packages/authorize/adapter"

	"github.com/casbin/casbin/v2"
)

type Roles []string
type Policy string

type Authorizer struct {
	*casbin.Enforcer
	mapRoleAndActions map[string][]string
}

func New() *Authorizer {
	a := adapter.NewAdapter(string(CommonPolicy))
	m, _ := model.NewModelFromString(Model)
	e, _ := casbin.NewEnforcer(m, a)
	return &Authorizer{
		Enforcer:          e,
		mapRoleAndActions: buildMapRoleActions(string(CommonPolicy)),
	}
}

// (Ex:
//
//	roles: ["admin", "shipper"]
//	action: "/CreateProduct:POST"
//
// )
func (a *Authorizer) Check(roles Roles, action string) bool {
	for _, role := range roles {
		res, err := a.Enforcer.Enforce(action, role)
		if err != nil {
			return false
		}
		return res
	}
	return false
}

func buildMapRoleActions(policy string) map[string][]string {
	lines := strings.Split(policy, "\n")
	checkedActions := map[string]struct{}{}

	// (Ex: m := make(map[string][]string)
	//		"admin": []string {
	// 			"/api/GetProduct:post"
	//			"/api/CreateProduct:post"
	//		},
	//		"shipper": []string {
	// 			"/api/GetProducts:post"
	//		}
	// )
	m := make(map[string][]string) // map role and actions
	for _, line := range lines {
		// prefix '#' for comment
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// p, /api/CreateProduct:POST, admin, shipper => ["p",    "/api/CreateProduct:POST", "admin", "shipper"]
		//                                                 ^              	   ^                	^
		//                                               prefix         	 action          	  roles......
		elements := strings.Split(line, ",")
		if len(elements) < 3 {
			panic(fmt.Sprintf("Invalid policy setup, error line content: %v", line))
		}

		action := strings.TrimSpace(elements[1])
		if _, ok := checkedActions[action]; ok {
			panic(fmt.Sprintf("Duplicate action, error line content: %v", line))
		}
		checkedActions[action] = struct{}{}

		roles := elements[2:]
		for _, role := range roles {
			role = strings.TrimSpace(role)
			if role == "" {
				panic(fmt.Sprintf("Invalid policy setup, error line content: %v", line))
			}
			m[role] = append(m[role], action)
		}
	}
	return m
}
