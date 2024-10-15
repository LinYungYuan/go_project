// internal/service/product_service.go
package service

import (
	"go_project/internal/model"
	"go_project/internal/repository"
	"time"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return s.productRepo.Create(product)
}

func (s *ProductService) GetProduct(id int64) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *ProductService) ListProducts() ([]*model.Product, error) {
	return s.productRepo.List()
}
