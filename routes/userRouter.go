package routes

import (
	controller "restro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/users", controller.GetUsers())
	incomingRouters.GET("/users/:user_id", controller.GetUser())
	incomingRouters.POST("/users/signup", controller.Signup())
	incomingRouters.POST("/users/logic", controller.Login())
}
