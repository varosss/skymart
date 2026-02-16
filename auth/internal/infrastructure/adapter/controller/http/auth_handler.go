package http

import (
	"errors"
	"net/http"

	"clirzy/auth/internal/application/usecase"
	"clirzy/auth/internal/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	registerUC     *usecase.RegisterUseCase
	loginUC        *usecase.LoginUseCase
	logoutUC       *usecase.LogoutUseCase
	refreshTokenUC *usecase.RefreshTokenUseCase
}

func NewAuthHandler(
	registerUC *usecase.RegisterUseCase,
	loginUC *usecase.LoginUseCase,
	logoutUC *usecase.LogoutUseCase,
	refreshTokenUC *usecase.RefreshTokenUseCase,
) *AuthHandler {
	return &AuthHandler{
		loginUC:        loginUC,
		logoutUC:       logoutUC,
		refreshTokenUC: refreshTokenUC,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.registerUC.Execute(
		c.Request.Context(),
		usecase.RegisterCommand{
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": result.UserID.String()})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.loginUC.Execute(
		c.Request.Context(),
		usecase.LoginCommand{
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := extractRefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = h.logoutUC.Execute(
		c.Request.Context(),
		usecase.LogoutCommand{
			RefreshToken: refreshToken,
		},
	)

	if err != nil {
		switch err {
		case domain.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := extractRefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	result, err := h.refreshTokenUC.Execute(
		c.Request.Context(),
		usecase.RefreshTokenCommand{
			RefreshToken: refreshToken,
		},
	)

	if err != nil {
		switch err {
		case domain.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"access_token":  result.AccessToken,
			"refresh_token": result.RefreshToken,
		},
	)
}

func extractRefreshToken(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil || cookie == "" {
		return "", errors.New("refresh token not found")
	}
	return cookie, nil
}
