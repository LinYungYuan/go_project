// internal/service/order_service.go
package service

import (
	"go_project/internal/model"
	"go_project/internal/repository"
	"time"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, productRepo *repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	for _, item := range order.Items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			return err
		}
		if product.Stock < item.Quantity {
			return err
		}
		item.Price = product.Price
		order.TotalPrice += item.Price * float64(item.Quantity)
	}

	order.Status = "pending"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	return s.orderRepo.Create(order)
}

func (s *OrderService) GetOrder(id int64, userID int64) (*model.Order, error) {
	return s.orderRepo.GetByID(id, userID)
}
