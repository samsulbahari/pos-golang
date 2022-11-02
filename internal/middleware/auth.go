package middleware

import (
	"context"
	"pos/internal/libraries"
	"strings"

	"github.com/gin-gonic/gin"
)

var bearer = "Bearer "

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"code":    "401",
				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}

		//untuk cek apakah header = bearer
		if !strings.HasPrefix(authHeader, bearer) {
			ctx.JSON(401, gin.H{
				"code":    "401",
				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}
		token := strings.Split(authHeader, " ")
		data, err := libraries.DecryptJwt(token[1])
		if err != nil {
			ctx.JSON(401, gin.H{
				"code":    "401",
				"message": "Invalid token",
			})

			ctx.Abort()
			return
		}
		userID := int(data["user_id"].(float64))
		roleID := int(data["role"].(float64))
		ctxUserID := context.WithValue(ctx.Request.Context(), "user_id", userID)
		ctxRole := context.WithValue(ctx.Request.Context(), "role", roleID)
		ctx.Request = ctx.Request.WithContext(ctxUserID)
		ctx.Request = ctx.Request.WithContext(ctxRole)

		ctx.Next()
	}

}
