package main

import (
	"fmt"
	"pos/internal/auth/handler_auth"
	"pos/internal/auth/repository_auth"
	"pos/internal/auth/service_auth"
	"pos/internal/category/handler_category"
	"pos/internal/category/repository_category"
	"pos/internal/category/service_category"
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
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Static("/assets", "./assets")
	r.Use(middleware.CORSMiddleware())

	db := database.ConnectDatabase()
	authRepo := repository_auth.NewAuthRepo(db)
	authService := service_auth.NewAuthService(authRepo)
	AuthHandler := handler_auth.NewAuthHandler(authService)

	auth := r.Group("auth")
	{
		auth.POST("/login", AuthHandler.Login)

	}

	menuRepo := repository_menu.NewMenuRepo(db)
	menuService := service_menu.NewMenuService(menuRepo)
	menuHandler := handler_menu.NewMenuHandler(menuService)

	categoryRepo := repository_category.NewCategoryRepo(db)
	categoryService := service_category.NewCategoryService(categoryRepo)
	categoryHandler := handler_category.NewCategoryHandler(categoryService)

	middleware := middleware.WithAuth()
	authorized := r.Group("")
	authorized.Use(middleware)
	{
		authorized.GET("/menu", menuHandler.GetMenu)
		authorized.GET("/category", categoryHandler.GetCategory)
		authorized.POST("/category", categoryHandler.CreateCategory)
	}

	r.Run()
}
