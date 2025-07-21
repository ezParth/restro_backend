package routes

import (
	controller "restro_backend/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/menus", controller.GetMenus())
	incomingRouters.GET("/menus/:menu_id", controller.GetMenu())
	incomingRouters.PUT("/menus", controller.CreateMenu())
	incomingRouters.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
