package v1

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/api/converter"
	authv1 "github.com/escoutdoor/linko/common/pkg/proto/auth/v1"
)

func (a *api) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	tokens, err := a.authService.Register(ctx, converter.ProtoRegisterRequestToCreateUserParams(req))
	if err != nil {
		return nil, err
	}

	return &authv1.RegisterResponse{
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	}, nil
}
