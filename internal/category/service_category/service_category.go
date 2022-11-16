package service_category

import (
	"errors"
	"math"
	"mime/multipart"
	"pos/internal/domain"
	"pos/internal/libraries"

	"github.com/gin-gonic/gin"
)

type CategoryRepository interface {
	GetData(offset int) ([]domain.MCategory, error)
	TotalData() (int64, error)
	Create(category domain.MCategory) error
}

type CategoryService struct {
	CategoriRepo CategoryRepository
}

func NewCategoryService(categoriRepo CategoryRepository) *CategoryService {
	return &CategoryService{categoriRepo}
}

func (cs CategoryService) GetDataService(page int) (domain.ResultCategory, error) {

	var Result domain.ResultCategory
	offset := (page - 1) * 10
	data, err := cs.CategoriRepo.GetData(offset)
	if err != nil {
		return Result, errors.New("Failed get data")
	}
	count, err := cs.CategoriRepo.TotalData()

	if err != nil {
		return Result, errors.New("Failed get data")
	}

	last_page_counts := float64(count) / float64(10)
	last_page := math.Ceil(last_page_counts)
	if last_page == 0 {
		last_page = 1
	}
	Result.Data = data
	Result.Page = page
	Result.PerPage = 10
	Result.Total = int(count)
	Result.LastPage = last_page

	return Result, nil
}

func (cs CategoryService) CreateCategoryService(file *multipart.FileHeader, ctx *gin.Context, category domain.MCategory, extension string) error {

	rand_string := libraries.RandomString()
	file_path := "assets/category/" + rand_string + "." + extension

	category.Image = ctx.Request.Host + "/" + file_path

	err := cs.CategoriRepo.Create(category)
	if err != nil {
		return errors.New("Failed get data")
	}

	ctx.SaveUploadedFile(file, file_path)
	return nil
}
