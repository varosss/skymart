package httptransport

import (
	"net/http"

	"clirzy/services/auth/internal/service"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service *service.AuthService
}

func NewHTTPHandler(service *service.AuthService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) Register(c *gin.Context) {
	var req AuthRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorAuthResponse{Error: err.Error()})

		return
	}

	output, err := h.service.Register(c.Request.Context(), AuthRequestToDomain(&req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorAuthResponse{Error: err.Error()})
	}

	resp := domainToRegisterResponse(output)

	c.JSON(http.StatusOK, resp)
}

func (h *HTTPHandler) Login(c *gin.Context) {
	var req AuthRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorAuthResponse{Error: err.Error()})

		return
	}

	output, err := h.service.Login(c.Request.Context(), AuthRequestToDomain(&req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorAuthResponse{Error: err.Error()})
	}

	resp := domainToLoginResponse(output)

	c.JSON(http.StatusOK, resp)
}
