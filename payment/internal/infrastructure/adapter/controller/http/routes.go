package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	handler *PaymentHandler,
) {
	payment := r.Group("/payment")
	{
		payment.POST("/yookassa", handler.HandleYookassaWebhook)
	}
}
