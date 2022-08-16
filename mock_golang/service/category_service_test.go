package service

import (
	"golang_tutorial_unit_test/mock_golang/entity"
	"golang_tutorial_unit_test/mock_golang/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	category1, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category1)
	assert.Equal(t, category.Id, category1.Id)
	assert.Equal(t, category.Name, category1.Name)
}
