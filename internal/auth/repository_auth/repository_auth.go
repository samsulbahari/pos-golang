package repository_auth

import (
	"pos/internal/domain"

	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (ar AuthRepo) GetUserByEmail(login domain.Login) (domain.User, error) {
	var user domain.User
	err := ar.db.Where("email = ?", login.Email).First(&user).Error
	return user, err
}
