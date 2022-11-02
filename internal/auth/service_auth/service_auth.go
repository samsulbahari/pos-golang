package service_auth

import (
	"errors"
	"pos/internal/domain"
	"pos/internal/libraries"
)

type AuthRepo interface {
	GetUserByEmail(login domain.Login) (domain.User, error)
}
type AuthService struct {
	AuthRepository AuthRepo
}

func NewAuthService(authrepo AuthRepo) *AuthService {
	return &AuthService{authrepo}
}

func (as AuthService) LoginService(login domain.Login) (string, error, int) {
	user, err := as.AuthRepository.GetUserByEmail(login)
	if err != nil {
		return "", errors.New("Failed Email or password"), 401
	}
	err = libraries.ComparePassword(user.Password, login.Password)
	if err != nil {
		return "", errors.New("Failed Email or password"), 401
	}

	jwt, err := libraries.GenerateJWT(user.ID, user.RoleId)
	if err != nil {
		return "", errors.New("Server Error generate jwt"), 500
	}
	return jwt, nil, 200

}
