package http

import (
	"clirzy/pkg/auth"
	aport "clirzy/product/internal/application/port"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	handler *ProductHandler,
	authClient aport.AuthGateway,
	sellers aport.SellerGateway,
	products aport.ProductQuery,
) {
	product := r.Group("/")
	product.Use(auth.Authenticate(authClient))
	{
		product.POST(
			"/products",
			handler.CreateProduct,
		)
		product.PATCH(
			"/products/:product_id",
			handler.UpdateProduct,
		)
	}
}
