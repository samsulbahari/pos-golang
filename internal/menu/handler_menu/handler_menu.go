package handler_menu

import (
	"pos/internal/domain"

	"github.com/gin-gonic/gin"
)

type MenuService interface {
	GetMenuService(role int) ([]domain.ResultMenu, error)
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
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
	return

}
