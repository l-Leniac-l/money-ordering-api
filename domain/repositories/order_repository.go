package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/l-leniac-l/money-ordering-api/domain/entities"
)

type OrderRepository interface {
	CreateOrder(order *entities.Order) *entities.Order
}

type OrderRepositoryImpl struct{}

func (or *OrderRepositoryImpl) CreateOrder(order *entities.Order) *entities.Order {
	id := uuid.New()

	newOrder := &entities.Order{
		OrderId:        id.String(),
		BankId:         order.BankId,
		Amount:         order.Amount,
		AdditionalInfo: order.AdditionalInfo,
	}

	fmt.Println("New order created", newOrder)

	return newOrder
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}
