package httptransport

import (
	"net/http"
	"strconv"

	"clirzy/services/user/internal/service"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service *service.UserService
}

func NewHTTPHandler(service *service.UserService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) CreateUsers(c *gin.Context) {
	var req CreateUsersBatchRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorUsersResponse{Error: err.Error()})

		return
	}

	users, err := h.service.CreateMany(c.Request.Context(), CreateUsersBatchRequestToDomain(req))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorUsersResponse{Error: err.Error()})

		return
	}

	createUsersResponse := DomainToCreateUsersBatchResponse(users)

	c.JSON(http.StatusOK, createUsersResponse)
}

func (h *HTTPHandler) GetUserById(c *gin.Context) {
	userIdToGetStr := c.Param("id")
	if userIdToGetStr == "" {
		c.JSON(http.StatusBadRequest, ErrorUsersResponse{Error: "id param is required"})

		return
	}

	userIdToGet, err := strconv.Atoi(userIdToGetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorUsersResponse{Error: "id param is invalid"})

		return
	}

	user, err := h.service.GetUserById(c.Request.Context(), userIdToGet)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorUsersResponse{Error: err.Error()})

		return
	}

	if user.ID != uint(c.GetInt("user_id")) {
		c.JSON(http.StatusBadRequest, ErrorUsersResponse{Error: "permission denied"})

		return
	}

	c.JSON(http.StatusBadRequest, DomainToUserResponse(user))
}
