package service

import (
	"fmt"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func CategoryServiceFactory() *CategoryService {
	return &CategoryService{repo: repository.NewCategoryRepository()}
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
	if category.Name == "" {
		return fmt.Errorf("category name is required")
	}
	if category.Slug == "" {
		return fmt.Errorf("category slug is required")
	}

	// Kiểm tra xem slug đã tồn tại chưa
	existingCategories, _, err := s.repo.FindAll(1, 1, category.Slug, true)
	if err != nil {
		return fmt.Errorf("error checking existing category: %v", err)
	}
	if len(existingCategories) > 0 {
		return fmt.Errorf("category with slug %s already exists", category.Slug)
	}

	return s.repo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) GetAllCategories(page, pageSize int, search string, status bool) ([]model.Category, int64, error) {
	return s.repo.FindAll(page, pageSize, search, status)
}

func (s *CategoryService) UpdateCategory(id uint, updatedCategory *model.Category) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if updatedCategory.Name != "" {
		category.Name = updatedCategory.Name
	}
	if updatedCategory.Slug != "" {
		// Kiểm tra xem slug mới đã tồn tại chưa
		existingCategories, _, err := s.repo.FindAll(1, 1, updatedCategory.Slug, true)
		if err != nil {
			return fmt.Errorf("error checking existing category: %v", err)
		}
		if len(existingCategories) > 0 && existingCategories[0].ID != id {
			return fmt.Errorf("category with slug %s already exists", updatedCategory.Slug)
		}
		category.Slug = updatedCategory.Slug
	}
	if updatedCategory.Description != "" {
		category.Description = updatedCategory.Description
	}
	if updatedCategory.ParentID != nil {
		category.ParentID = updatedCategory.ParentID
	}

	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
