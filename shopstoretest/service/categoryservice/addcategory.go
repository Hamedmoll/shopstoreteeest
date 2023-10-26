package categoryservice

import (
	"fmt"
	"shopstoretest/entity"
	"shopstoretest/param"
)

func (s Service) AddCategory(req param.AddCategoryRequest) (param.AddCategoryResponse, error) {
	exist, eErr := s.Repository.CheckExistCategory(req.Name)
	if eErr != nil {

		return param.AddCategoryResponse{}, fmt.Errorf("unexpected error %w", eErr)
	}

	if exist {

		return param.AddCategoryResponse{}, fmt.Errorf("category is exist already")
	}

	newCategory := entity.Category{
		ID:   0,
		Name: req.Name,
	}

	createdCategory, aErr := s.Repository.AddCategory(newCategory)
	if aErr != nil {

		return param.AddCategoryResponse{}, fmt.Errorf("unexpected error %w", aErr)
	}

	categoryInfo := param.CategoryInfo{Name: createdCategory.Name}

	res := param.AddCategoryResponse{
		CategoryInfo: categoryInfo,
	}

	return res, nil
}
