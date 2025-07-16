package service

import (
	"go_ecommerce/internal/dto"
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

// #region customization
func (s *ProductService) SearchForCard(name string, page, pageSize int) ([]dto.ProductCardResponse, int64, error) {
	products, total, err := s.repo.FindForCard(name, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	// filter out-of-stock products
	filtered := make([]dto.ProductCardResponse, 0)
	for _, p := range products {
		if p.Status == model.ProductStatusActive {
			filtered = append(filtered, p)
		}
	}
	return filtered, total, nil
}
