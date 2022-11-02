package main

import (
	"fmt"
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
	database.ConnectDatabase()

	r := gin.Default()

	db := database.ConnectDatabase()
	authRepo := repository_auth.NewAuthRepo(db)
	authService := service_auth.NewAuthService(authRepo)
	AuthHandler := handler_auth.NewAuthHandler(authService)

	middleware := middleware.WithAuth()

	auth := r.Group("/auth")
	{
		auth.POST("/login", AuthHandler.Login)

	}

	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware)
	{
		menuRepo := repository_menu.NewMenuRepo(db)
		menuService := service_menu.NewMenuService(menuRepo)
		menuHandler := handler_menu.NewMenuHandler(menuService)
		authorized.GET("/menu", menuHandler.GetMenu)
	}

	r.Run()
}
