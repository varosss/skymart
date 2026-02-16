package http

import (
	aport "clirzy/order/internal/application/port"
	"clirzy/pkg/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	handler *OrderHandler,
	authClient aport.AuthGateway,
) {
	order := r.Group("/")
	order.Use(auth.Authenticate(authClient))
	{
		order.POST("/orders", handler.CreateOrder)
	}
}
