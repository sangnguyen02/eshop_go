package service

import (
	"fmt"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/repository"
)

type BrandService struct {
	repo *repository.BrandRepository
}

func BrandServiceFactory() *BrandService {
	return &BrandService{repo: repository.NewBrandRepository()}
}

func (s *BrandService) CreateBrand(brand *model.Brand) error {
	if brand.Name == "" {
		return fmt.Errorf("brand name is required")
	}
	if brand.Slug == "" {
		return fmt.Errorf("brand slug is required")
	}

	// Kiểm tra xem slug đã tồn tại chưa
	existingBrand, _, err := s.repo.FindAll(1, 1, brand.Slug, true)
	if err != nil {
		return fmt.Errorf("error checking existing category: %v", err)
	}
	if len(existingBrand) > 0 {
		return fmt.Errorf("category with slug %s already exists", brand.Slug)
	}

	return s.repo.Create(brand)
}

func (s *BrandService) GetBrandByID(id uint) (*model.Brand, error) {
	return s.repo.FindByID(id)
}

func (s *BrandService) GetAllBrands(page, pageSize int, search string, status bool) ([]model.Brand, int64, error) {
	return s.repo.FindAll(page, pageSize, search, status)
}

func (s *BrandService) UpdateBrand(id uint, updatedBrand *model.Brand) error {
	brand, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if updatedBrand.Name != "" {
		brand.Name = updatedBrand.Name
	}

	if updatedBrand.Slug != "" {
		// Kiểm tra xem slug mới đã tồn tại chưa
		existingBrand, _, err := s.repo.FindAll(1, 1, updatedBrand.Slug, true)
		if err != nil {
			return fmt.Errorf("error checking existing category: %v", err)
		}
		if len(existingBrand) > 0 && existingBrand[0].ID != id {
			return fmt.Errorf("category with slug %s already exists", updatedBrand.Slug)
		}
		brand.Slug = updatedBrand.Slug
	}

	if updatedBrand.Description != "" {
		brand.Description = updatedBrand.Description
	}

	if updatedBrand.LogoURL != "" {
		brand.LogoURL = updatedBrand.LogoURL
	}

	brand.Status = updatedBrand.Status

	return s.repo.Update(brand)
}

func (s *BrandService) DeleteBrand(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
