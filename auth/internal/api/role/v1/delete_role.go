package v1

import (
	"context"

	rolev1 "github.com/escoutdoor/linko/common/pkg/proto/role/v1"
)

func (a *api) DeleteRole(ctx context.Context, req *rolev1.DeleteRoleRequest) (*rolev1.DeleteRoleResponse, error) {
	if err := a.roleService.DeleteRole(ctx, req.GetRoleId()); err != nil {
		return nil, err
	}

	return &rolev1.DeleteRoleResponse{}, nil
}
