package handler_category

import (
	"pos/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	GetDataService(page int) (domain.ResultCategory, error)
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
		ctx.JSON(400, gin.H{
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
