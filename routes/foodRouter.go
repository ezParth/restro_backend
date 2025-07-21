package routes

import (
	controller "restro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/foods", controller.GetFoods())
	incomingRouters.GET("/foods/:user_id", controller.GetFood())
	incomingRouters.POST("/foods", controller.CreateFood())
	incomingRouters.PATCH("/foods/:food_id", controller.UpdateFood())
}
