package user

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/entity"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) GetUser(ctx context.Context, userID string) (entity.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return entity.User{}, errwrap.Wrap("get user from repository", err)
	}

	return user, nil
}
