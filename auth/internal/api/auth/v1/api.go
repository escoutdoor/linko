package v1

import (
	"github.com/escoutdoor/linko/auth/internal/service"
	authv1 "github.com/escoutdoor/linko/common/pkg/proto/auth/v1"
)

type api struct {
	authService service.AuthService
	authv1.UnimplementedAuthServiceServer
}

func NewImplementation(authService service.AuthService) *api {
	return &api{
		authService: authService,
	}
}
