package routes

import (
	controller "restro_backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/orderItems", controller.GetOrderItems())
	incomingRouters.GET("/orderItems/:orderItem_id", controller.GetOrderItems())
	incomingRouters.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRouters.PUT("/orderItems", controller.CreateOrderItems())
	incomingRouters.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItems())
}
