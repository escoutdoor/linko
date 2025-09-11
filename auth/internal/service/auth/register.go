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

	var userID string
	if txErr := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error

		userID, err = s.userRepository.CreateUser(ctx, in)
		if err != nil {
			return errwrap.Wrap("create user", err)
		}

		if err := s.userRoleRepository.AddRolesToUser(ctx, userID, in.Roles...); err != nil {
			return errwrap.Wrap("add roles to user", err)
		}

		return nil
	}); txErr != nil {
		return entity.TokenPair{}, txErr
	}

	tokens, err := s.tokenProvider.GenerateTokens(userID)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("generate jwt tokens", err)
	}

	return tokens, nil
}
