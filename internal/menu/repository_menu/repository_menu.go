package repository_menu

import (
	"pos/internal/domain"

	"gorm.io/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *MenuRepo {
	return &MenuRepo{db}
}

func (mr MenuRepo) Getdata(role int) ([]domain.Menu, error) {
	var menu []domain.Menu

	err := mr.db.Table("x_role_has_permission").Joins("JOIN r_menu on x_role_has_permission.menu_id = r_menu.id").Select("r_menu.name", "x_role_has_permission.id").Order("menu_id asc").Where("role_id = ?", role).Scan(&menu).Error

	return menu, err
}

func (mr MenuRepo) GetSubmenu(idpermission int) ([]domain.Submenu, error) {
	var submenu []domain.Submenu
	//err := mr.db.Preload("Submenu.Menu").Where("role_id = ?", "1").Order("r_menu.id asc").Joins("JOIN x_role_has_permission on x_role_has_permission.menu_id = r_menu.id").Find(&menu).Error

	err := mr.db.Table("x_role_has_permission_submenu").Joins("JOIN r_submenu on x_role_has_permission_submenu.submenu_id = r_submenu.id").Select("r_submenu.name").Order("r_submenu.id asc").Where("permission_id = ?", idpermission).Scan(&submenu).Error

	return submenu, err
}
