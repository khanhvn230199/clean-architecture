package handler

import (
	"github.com/example-golang-projects/clean-architecture/services/user/dependency"
	"github.com/example-golang-projects/clean-architecture/services/user/modules/role/usecase"
	"net/http"
)

type RoleHandler struct {
	dependency  dependency.UserDependency
	roleUseCase usecase.RoleUseCase
}

func NewRoleHandler(userDependency dependency.UserDependency) RoleHandler {
	return RoleHandler{
		dependency:  userDependency,
		roleUseCase: usecase.NewRoleUseCase(userDependency.DB),
	}
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	h.roleUseCase.CreateRole(nil, 31)
}
