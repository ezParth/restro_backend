package routes

import (
	controller "restro_backend/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/tables", controller.GetTables())
	incomingRouters.GET("/tables/:table_id", controller.GetTable())
	incomingRouters.PUT("/tables", controller.CreateTable())
	incomingRouters.PATCH("/tables/:table_id", controller.UpdateTable())
}
