package service_category

import (
	"errors"
	"math"
	"mime/multipart"
	"os"
	"pos/internal/domain"
	"pos/internal/libraries"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryRepository interface {
	GetData(offset int) ([]domain.MCategory, error)
	TotalData() (int64, error)
	Create(category domain.MCategory) error
	FindByid(id int) (domain.MCategory, error)
	Delete(id int) error
	Update(id int, category domain.Update) error
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
func (cs CategoryService) DeleteService(id int) error {
	_, err := cs.CategoriRepo.FindByid(id)
	if err != nil {
		return errors.New("Id not found")
	}
	err = cs.CategoriRepo.Delete(id)
	if err != nil {
		return errors.New("Failed delete data")
	}
	return nil
}

func (cs CategoryService) UpdateService(id int, file *multipart.FileHeader, ctx *gin.Context, category domain.Update, extension string) error {
	data, err := cs.CategoriRepo.FindByid(id)
	if err != nil {
		return errors.New("Id not found")
	}

	lastimage_path := strings.Split(data.Image, "/")

	err = os.Remove("assets/category/" + lastimage_path[3])

	if err != nil {
		return errors.New("Error remove last image")
	}

	rand_string := libraries.RandomString()
	file_path := "assets/category/" + rand_string + "." + extension

	category.Image = ctx.Request.Host + "/" + file_path

	err = cs.CategoriRepo.Update(id, category)
	if err != nil {
		return errors.New("Error update data")
	}
	ctx.SaveUploadedFile(file, file_path)
	return nil
}
