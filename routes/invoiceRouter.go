package routes

import (
	controller "restro_backend/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRouters *gin.Engine) {
	incomingRouters.GET("/invoices", controller.GetInvoices())
	incomingRouters.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRouters.PUT("/invoices", controller.CreateInvoice())
	incomingRouters.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
