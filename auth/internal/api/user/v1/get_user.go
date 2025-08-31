package v1

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/api/converter"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
)

func (a *api) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := a.userService.GetUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	protoUser := converter.UserToProtoUser(user)
	return &userv1.GetUserResponse{User: protoUser}, nil
}
