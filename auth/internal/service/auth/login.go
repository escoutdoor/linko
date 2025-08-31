package auth

import (
	"context"
	"errors"

	"github.com/escoutdoor/linko/auth/internal/entity"
	apperrors "github.com/escoutdoor/linko/auth/internal/errors"
	"github.com/escoutdoor/linko/auth/internal/errors/codes"
	"github.com/escoutdoor/linko/auth/internal/utils/hasher"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) Login(ctx context.Context, email, password string) (entity.TokenPair, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) && appErr.Code == codes.UserNotFound {
			return entity.TokenPair{}, apperrors.ErrIncorrectCreadentials
		}

		return entity.TokenPair{}, errwrap.Wrap("get user by email from repository", err)
	}

	if match := hasher.CompareHashAndPassword(user.Password, password); !match {
		return entity.TokenPair{}, apperrors.ErrIncorrectCreadentials
	}

	tokens, err := s.tokenProvider.GenerateTokens(user.ID)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("generate jwt tokens", err)
	}

	return tokens, nil
}
