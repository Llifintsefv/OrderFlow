package service

import (
	"order-service/internal/repository"
	"order-service/internal/models"
)

type OrderService interface {
	CreateOrder(order *models.Order) error
	GetOrder(OrderId int) (*models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order *models.Order) error {
	///////////////////////////////
	return s.repo.CreateOrder(order)
}

func (s *orderService) GetOrder(OrderId int) (*models.Order, error) {
	return s.repo.GetOrder(OrderId)
}
