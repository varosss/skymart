package httptransport

import (
	"clirzy/auth/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service *service.AuthService
}

func NewHTTPHandler(service *service.AuthService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"authToken": "fadsfadf"})
}

func (h *HTTPHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"authToken": "fadsfadf"})
}
