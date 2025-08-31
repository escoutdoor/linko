package auth

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/entity"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) RefreshToken(ctx context.Context, refreshToken string) (entity.TokenPair, error) {
	userID, err := s.tokenProvider.ValidateRefreshToken(refreshToken)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("validate refresh token", err)
	}

	if _, err := s.userRepository.GetUserByID(ctx, userID); err != nil {
		return entity.TokenPair{}, errwrap.Wrap("get user by if from repository", err)
	}

	tokens, err := s.tokenProvider.GenerateTokens(userID)
	if err != nil {
		return entity.TokenPair{}, errwrap.Wrap("generate jwt tokens", err)
	}

	return tokens, nil
}
