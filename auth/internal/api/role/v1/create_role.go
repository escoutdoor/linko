package v1

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/api/converter"
	rolev1 "github.com/escoutdoor/linko/common/pkg/proto/role/v1"
)

func (a *api) CreateRole(ctx context.Context, req *rolev1.CreateRoleRequest) (*rolev1.CreateRoleResponse, error) {
	role, err := a.roleService.CreateRole(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	protoRole := converter.RoleToProtoRole(role)
	return &rolev1.CreateRoleResponse{Role: protoRole}, nil
}
