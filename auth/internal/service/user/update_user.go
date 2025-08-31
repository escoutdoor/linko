package user

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) UpdateUser(ctx context.Context, in dto.UpdateUserParams) (entity.User, error) {
	if _, err := s.userRepository.GetUserByID(ctx, in.ID); err != nil {
		return entity.User{}, errwrap.Wrap("get user from repository", err)
	}

	if err := s.userRepository.UpdateUser(ctx, in); err != nil {
		return entity.User{}, errwrap.Wrap("update user", err)
	}

	user, err := s.userRepository.GetUserByID(ctx, in.ID)
	if err != nil {
		return entity.User{}, errwrap.Wrap("get updated user from repository", err)
	}

	return user, nil
}
