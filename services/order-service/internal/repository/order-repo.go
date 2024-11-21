package repository

import (
	"database/sql"
	"order-service/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrder(OrderId int) (*models.Order,error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	db.AutoMigrate(&models.Order{}, &models.Item{})
	return &orderRepository{db: db}
}

func (o *orderRepository) CreateOrder(order *models.Order) error {
	err := o.db.Create(order).Error
	return err
}

func (o *orderRepository) GetOrder(OrderId int) (*models.Order, error) {
	order := &models.Order{}

	err := o.db.Where("id = ?", OrderId).First(order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return order, nil
	
}
