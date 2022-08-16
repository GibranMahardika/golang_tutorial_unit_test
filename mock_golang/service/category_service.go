package service

import (
	"errors"
	"golang_tutorial_unit_test/mock_golang/entity"
	"golang_tutorial_unit_test/mock_golang/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
