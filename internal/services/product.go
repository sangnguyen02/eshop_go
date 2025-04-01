package services

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) SearchProducts(name string) ([]models.Product, error) {
	return s.repo.FindByName(name)
}
