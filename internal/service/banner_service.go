package service

import (
	"fmt"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/repository"
)

type BannerService struct {
	repo *repository.BannerRepository
}

func BannerServiceFactory() *BannerService {
	return &BannerService{repo: repository.NewBannerRepository()}
}

func (s *BannerService) CreateBanner(banner *model.Banner) error {
	if banner.Name == "" {
		return fmt.Errorf("banner name is required")
	}

	if banner.Image == "" {
		return fmt.Errorf("banner image is required")
	}

	return s.repo.Create(banner)
}

func (s *BannerService) GetBannerByID(id uint) (*model.Banner, error) {
	return s.repo.FindByID(id)
}

func (s *BannerService) GetAllBanners(page, pageSize int, search string, status bool) ([]model.Banner, int64, error) {
	return s.repo.FindAll(page, pageSize, search, status)
}

func (s *BannerService) UpdateBanner(id uint, updatedBanner *model.Banner) error {
	banner, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if updatedBanner.Name != "" {
		banner.Name = updatedBanner.Name
	}

	if updatedBanner.Image != "" {
		banner.Image = updatedBanner.Image
	}

	banner.Status = updatedBanner.Status

	return s.repo.Update(banner)
}

func (s *BannerService) DeleteBanner(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
