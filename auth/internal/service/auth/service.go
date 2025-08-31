package auth

import (
	"github.com/escoutdoor/linko/auth/internal/repository"
	"github.com/escoutdoor/linko/auth/internal/utils/token"
)

type service struct {
	userRepository repository.UserRepository
	tokenProvider  token.Provider
}

func NewService(userRepository repository.UserRepository, tokenProvider token.Provider) *service {
	return &service{
		userRepository: userRepository,
		tokenProvider:  tokenProvider,
	}
}
