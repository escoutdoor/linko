package v1

import (
	"context"

	// "github.com/escoutdoor/linko/auth/internal/api/converter"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
)

func (a *api) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	// if err := i.userService.UpdateUser(ctx, ); err != nil {
	// 	return nil, err
	// }

	return &userv1.UpdateUserResponse{}, nil
}
