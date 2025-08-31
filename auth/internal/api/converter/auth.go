package converter

import (
	"github.com/escoutdoor/linko/auth/internal/dto"
	authv1 "github.com/escoutdoor/linko/common/pkg/proto/auth/v1"
)

func ProtoRegisterRequestToCreateUserParams(req *authv1.RegisterRequest) dto.CreateUserParams {
	return dto.CreateUserParams{
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		Email:       req.GetEmail(),
		PhoneNumber: req.GetPhoneNumber(),
		Password:    req.GetPassword(),
	}
}
