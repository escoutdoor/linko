package converter

import (
	"github.com/escoutdoor/linko/auth/internal/entity"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProtoUser(user entity.User) *userv1.User {
	return &userv1.User{
		Id:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
	}
}

func UsersToProtoUsers(users []entity.User) []*userv1.User {
	list := make([]*userv1.User, 0, len(users))
	for _, u := range users {
		list = append(list, UserToProtoUser(u))
	}

	return list
}
