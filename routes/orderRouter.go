package routes

import (
	controller "restro_backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/orders", controller.GetOrders())
	incomingRouters.GET("/orders/:order_id", controller.GetOrder())
	incomingRouters.PUT("/orders", controller.CreateOrder())
	incomingRouters.PATCH("/orders/:order_id", controller.UpdateOrder())
}
