package repositories

import (
	"testing"

	"github.com/l-leniac-l/money-ordering-api/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	order := &entities.Order{
		BankId:         "f0b48033-dbb2-4bd5-b24d-8f8763e9461f",
		Amount:         100000.56,
		AdditionalInfo: "credit cards",
	}

	or := NewOrderRepository()

	newOrder := or.CreateOrder(order)

	assert.IsType(t, &entities.Order{}, newOrder)
	assert.Equal(t, newOrder.Amount, order.Amount)
	assert.Equal(t, newOrder.BankId, order.BankId)
}
