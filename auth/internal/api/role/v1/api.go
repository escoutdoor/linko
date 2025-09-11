package v1

import (
	"github.com/escoutdoor/linko/auth/internal/service"
	rolev1 "github.com/escoutdoor/linko/common/pkg/proto/role/v1"
)

type api struct {
	roleService service.RoleService
	rolev1.UnimplementedRoleServiceServer
}

func NewImplementation(roleService service.RoleService) *api {
	return &api{
		roleService: roleService,
	}
}
