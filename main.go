package main

import (
	"fmt"
	"net/http"
	"pos/internal/auth/handler_auth"
	"pos/internal/auth/repository_auth"
	"pos/internal/auth/service_auth"
	"pos/internal/database"
	"pos/internal/menu/handler_menu"
	"pos/internal/menu/repository_menu"
	"pos/internal/menu/service_menu"
	"pos/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	db := database.ConnectDatabase()
	authRepo := repository_auth.NewAuthRepo(db)
	authService := service_auth.NewAuthService(authRepo)
	AuthHandler := handler_auth.NewAuthHandler(authService)

	auth := r.Group("auth")
	{
		auth.POST("/login", AuthHandler.Login)

	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
		})
	})

	menuRepo := repository_menu.NewMenuRepo(db)
	menuService := service_menu.NewMenuService(menuRepo)
	menuHandler := handler_menu.NewMenuHandler(menuService)

	middleware := middleware.WithAuth()
	authorized := r.Group("")
	authorized.Use(middleware)
	{
		authorized.GET("/menu", menuHandler.GetMenu)
	}

	r.Run()
}
