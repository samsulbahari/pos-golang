package service_menu

import (
	"errors"
	"pos/internal/domain"
)

type MenuRepo interface {
	Getdata(role int) ([]domain.Menu, error)
	GetSubmenu(idpermission int) ([]domain.Submenu, error)
}

type MenuService struct {
	MenuRepository MenuRepo
}

func NewMenuService(Mr MenuRepo) *MenuService {
	return &MenuService{Mr}
}

func (ms MenuService) GetMenuService(role int) ([]domain.Result, error) {
	data, err := ms.MenuRepository.Getdata(role)
	if err != nil {
		return nil, errors.New("failed get data")
	}

	result := make([]domain.Result, 0)

	for _, getdata := range data {
		submenu, _ := ms.MenuRepository.GetSubmenu(getdata.ID)
		result = append(result, domain.Result{
			ID:      getdata.ID,
			Name:    getdata.Name,
			Submenu: submenu,
		})
	}

	return result, nil

}
