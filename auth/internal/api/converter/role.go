package converter

import (
	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	rolev1 "github.com/escoutdoor/linko/common/pkg/proto/role/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtoListRolesRequestToListRolesParams(req *rolev1.ListRolesRequest) dto.ListRolesParams {
	return dto.ListRolesParams{
		RoleIDs: req.GetRoleIds(),
	}
}

func RoleToProtoRole(role entity.Role) *rolev1.Role {
	return &rolev1.Role{
		Id:        role.ID,
		Name:      role.Name,
		CreatedAt: timestamppb.New(role.CreatedAt),
		UpdatedAt: timestamppb.New(role.UpdatedAt),
	}
}

func RolesToProtoRoles(roles []entity.Role) []*rolev1.Role {
	list := make([]*rolev1.Role, 0, len(roles))
	for _, r := range roles {
		list = append(list, RoleToProtoRole(r))
	}

	return list
}
