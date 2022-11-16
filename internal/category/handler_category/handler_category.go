package handler_category

import (
	"mime/multipart"
	"pos/internal/domain"
	"pos/internal/libraries"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	GetDataService(page int) (domain.ResultCategory, error)
	CreateCategoryService(file *multipart.FileHeader, ctx *gin.Context, category domain.MCategory, extension string) error
}

type CategoryHandler struct {
	CategoryServ CategoryService
}

func NewCategoryHandler(CategoryServ CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryServ}
}

func (ch CategoryHandler) GetCategory(ctx *gin.Context) {
	pageParam := ctx.Query("page")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		ctx.JSON(422, gin.H{
			"message": "Invalid input ID",
		})
		return
	}
	data, err := ch.CategoryServ.GetDataService(page)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, data)
	return

}
func (ch CategoryHandler) CreateCategory(ctx *gin.Context) {
	var category domain.MCategory
	err := ctx.ShouldBind(&category)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(422, gin.H{
			"message": validation_response,
		})
		return
	}
	file, _ := ctx.FormFile("image")
	if file == nil {
		ctx.JSON(422, gin.H{
			"message": "Image cant null",
		})
		return
	}
	imageExtension := strings.Split(file.Filename, ".")

	if imageExtension[1] != "jpg" && imageExtension[1] != "jpeg" && imageExtension[1] != "png" {

		ctx.JSON(422, gin.H{
			"message": "Image must jpg,png,jpeg",
		})
		return
	}

	err = ch.CategoryServ.CreateCategoryService(file, ctx, category, imageExtension[1])
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"message": "Success create data",
	})
	return

}
