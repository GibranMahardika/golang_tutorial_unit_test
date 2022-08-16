package repository

import (
	"golang_tutorial_unit_test/mock_golang/entity"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
