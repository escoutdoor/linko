package v1

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/api/converter"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
)

func (a *api) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	params, err := converter.ProtoUpdateUserRequestToUpdateUserParams(req)
	if err != nil {
		return nil, err
	}

	user, err := a.userService.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	protoUser := converter.UserToProtoUser(user)
	return &userv1.UpdateUserResponse{User: protoUser}, nil
}
