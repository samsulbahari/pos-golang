package handler_auth

import (
	"pos/internal/domain"
	"pos/internal/libraries"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginService(login domain.Login) (string, error, int)
}
type AuthHandler struct {
	authservice AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (ah AuthHandler) Login(ctx *gin.Context) {
	var login domain.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(422, gin.H{
			"message": validation_response,
		})
		return
	}
	user, err, code := ah.authservice.LoginService(login)
	if err != nil {
		ctx.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"message": "Success Login",
		"token":   user,
	})

}
