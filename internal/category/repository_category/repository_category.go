package repository_category

import (
	"pos/internal/domain"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

func (cr CategoryRepo) GetData(offset int) ([]domain.MCategory, error) {
	var category []domain.MCategory
	err := cr.db.Limit(10).Offset(offset).Find(&category).Error

	return category, err
}

func (cr CategoryRepo) TotalData() (int64, error) {
	var count int64
	var category []domain.MCategory
	err := cr.db.Model(&category).Count(&count).Error

	return count, err
}
