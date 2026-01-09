package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.Engine,
	handler *AuthHandler,
) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/refresh", handler.RefreshToken)
	}
}
