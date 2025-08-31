package auth

import (
	"context"
	"errors"

	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	apperrors "github.com/escoutdoor/linko/auth/internal/errors"
	"github.com/escoutdoor/linko/auth/internal/errors/codes"
	"github.com/escoutdoor/linko/auth/internal/utils/hasher"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) Register(ctx context.Context, in dto.CreateUserParams) (entity.TokenPair, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, in.Email)
	if err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			if appErr.Code != codes.UserNotFound {
				return entity.TokenPair{}, errwrap.Wrap("get user by email", err)
			}
		} else {
			return entity.TokenPair{}, errwrap.Wrap("get user by email", err)
		}
	}
	if user.Email != "" {
		return entity.TokenPair{}, apperrors.EmailAlreadyExists(in.Email)
	}

	pw, err := hasher.HashPassword(in.Password)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("hash password", err)
	}
	in.Password = pw

	userID, err := s.userRepository.CreateUser(ctx, in)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("create user", err)
	}

	tokens, err := s.tokenProvider.GenerateTokens(userID)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("generate jwt tokens", err)
	}

	return tokens, nil
}
