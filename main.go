package main

import (
	"fmt"
	"log"
	"os"
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

	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
)

func main() {

	if err := os.MkdirAll("assets/category", os.ModePerm); err != nil {
		fmt.Println("Folder already exits")
	}

	os.Mkdir("assets", 0700)

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

	store := persistence.NewInMemoryStore(time.Second)

	authorized.Use(middleware)
	{
		authorized.GET("/menu", cache.CachePage(store, time.Minute, menuHandler.GetMenu))
		authorized.GET("/category", categoryHandler.GetCategory)
		authorized.POST("/category", categoryHandler.CreateCategory)
		authorized.DELETE("/category", categoryHandler.DeleteCategory)
		authorized.PUT("/category", categoryHandler.UpdateCategory)
		authorized.PATCH("/category", categoryHandler.UpdateCategory)

	}

	r.Run()
}
