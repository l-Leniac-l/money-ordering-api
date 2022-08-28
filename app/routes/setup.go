package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/l-leniac-l/money-ordering-api/domain/repositories"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	orderRepository := repositories.NewOrderRepository()

	orders := router.Group("/orders")
	{
		orders.POST("", func(ctx *gin.Context) {
			CreateOrderHandler(ctx, orderRepository)
		})
	}
	router.GET("/health", HealthCheckHandler)

	return router
}
