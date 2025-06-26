package service

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func ProductServiceFactory() *ProductService {
	return &ProductService{repo: repository.NewProductRepository()}
}

func (s *ProductService) SearchProducts(name string, page, pageSize int) ([]model.Product, int64, error) {
	return s.repo.FindByName(name, page, pageSize)
}

func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
