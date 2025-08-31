package v1

import (
	"github.com/escoutdoor/linko/auth/internal/service"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
)

type api struct {
	userService service.UserService
	userv1.UnimplementedUserServiceServer
}

func NewImplementation(userService service.UserService) *api {
	return &api{
		userService: userService,
	}
}
