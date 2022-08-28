package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/l-leniac-l/money-ordering-api/domain/entities"
	"github.com/l-leniac-l/money-ordering-api/domain/repositories"
)

func CreateOrderHandler(c *gin.Context, r repositories.OrderRepository) {
	var order entities.Order

	if err := c.ShouldBindBodyWith(&order, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	newOrder := r.CreateOrder(&order)

	c.JSON(http.StatusCreated, newOrder)
}
