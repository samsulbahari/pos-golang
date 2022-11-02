package handler_menu

import (
	"pos/internal/domain"

	"github.com/gin-gonic/gin"
)

type MenuService interface {
	GetMenuService(role int) ([]domain.Result, error)
}

type MenuHandler struct {
	menuService MenuService
}

func NewMenuHandler(ms MenuService) *MenuHandler {
	return &MenuHandler{ms}
}

func (mh MenuHandler) GetMenu(ctx *gin.Context) {

	roleID := ctx.Request.Context().Value("role").(int)

	data, err := mh.menuService.GetMenuService(roleID)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
	return

}
