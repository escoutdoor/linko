package v1

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/api/converter"
	rolev1 "github.com/escoutdoor/linko/common/pkg/proto/role/v1"
)

func (a *api) ListRoles(ctx context.Context, req *rolev1.ListRolesRequest) (*rolev1.ListRolesResponse, error) {
	roles, err := a.roleService.ListRoles(ctx, converter.ProtoListRolesRequestToListRolesParams(req))
	if err != nil {
		return nil, err
	}

	protoRoles := converter.RolesToProtoRoles(roles)
	return &rolev1.ListRolesResponse{Roles: protoRoles}, nil
}
