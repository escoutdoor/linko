package converter

import (
	"fmt"

	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	apperrors "github.com/escoutdoor/linko/auth/internal/errors"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
	"github.com/nyaruka/phonenumbers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	firstNameFieldName   = "first_name"
	lastNameFieldName    = "last_name"
	emailFieldName       = "email"
	phoneNumberFieldName = "phone_number"
	passwordFieldName    = "password"
	rolesFieldName       = "roles"
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

func ProtoUpdateUserRequestToUpdateUserParams(req *userv1.UpdateUserRequest) (dto.UpdateUserParams, error) {
	params := dto.UpdateUserParams{
		ID: req.GetUserId(),
	}

	mask := req.GetUpdateMask()
	if mask == nil || len(mask.GetPaths()) == 0 {
		return dto.UpdateUserParams{}, apperrors.ValidationFailed("update mask is not provided or empty")
	}

	update := req.GetUpdate()
	if update == nil {
		return dto.UpdateUserParams{}, apperrors.ValidationFailed("update body is missing while update mask is provided")
	}

	for _, path := range mask.GetPaths() {
		switch path {
		case firstNameFieldName:
			v := update.GetFirstName()
			if v == "" {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed("first name is specified in update mask, but has no value")
			}
			params.FirstName = &v

		case lastNameFieldName:
			v := update.GetLastName()
			if v == "" {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed("last name is specified in update mask, but has no value")
			}
			params.LastName = &v

		case emailFieldName:
			v := update.GetEmail()
			if v == "" {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed("email is specified in update mask, but has no value")
			}
			params.Email = &v

		case phoneNumberFieldName:
			v := update.GetPhoneNumber()
			if v == "" {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed("phone number is specified in update mask, but has no value")
			}

			if err := validatePhoneNumber(v); err != nil {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed(err.Error())
			}
			params.PhoneNumber = &v

		case passwordFieldName:
			v := update.GetPassword()
			if v == "" {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed("password is specified in update mask, but has no value")
			}
			params.Password = &v

		case rolesFieldName:
			v := update.GetRoles()
			if len(v) == 0 {
				return dto.UpdateUserParams{}, apperrors.ValidationFailed("roles are specified in update mask, but have no value")
			}
			params.Roles = append([]string(nil), v...)

		}
	}

	return params, nil
}

func validatePhoneNumber(phoneNumber string) error {
	num, err := phonenumbers.Parse(phoneNumber, "UA")
	if err != nil {
		return fmt.Errorf("invalid phone number")
	}

	if !phonenumbers.IsValidNumberForRegion(num, "UA") {
		return fmt.Errorf("invalid phone number for your region (Ukraine)")
	}

	return nil
}
