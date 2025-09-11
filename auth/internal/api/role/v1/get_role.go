package v1

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/api/converter"
	rolev1 "github.com/escoutdoor/linko/common/pkg/proto/role/v1"
)

func (a *api) GetRole(ctx context.Context, req *rolev1.GetRoleRequest) (*rolev1.GetRoleResponse, error) {
	role, err := a.roleService.GetRole(ctx, req.GetRoleId())
	if err != nil {
		return nil, err
	}

	protoRole := converter.RoleToProtoRole(role)
	return &rolev1.GetRoleResponse{Role: protoRole}, nil
}
