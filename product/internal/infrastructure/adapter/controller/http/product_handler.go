package http

import (
	"clirzy/product/internal/application/usecase"
	"clirzy/product/internal/domain/valueobject"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	updateProductUC *usecase.UpdateProductUseCase
	createProductUC *usecase.CreateProductUseCase
}

func NewProductHandler(
	createProductUC *usecase.CreateProductUseCase,
	updateProductUC *usecase.UpdateProductUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUC: createProductUC,
		updateProductUC: updateProductUC,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
		Currency    string `json:"currency"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := valueobject.ParseUserID(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	money, err := valueobject.NewMoney(req.Price, req.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID, err := h.createProductUC.Execute(c.Request.Context(), usecase.CreateProductCommand{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Money:       money,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID.String()})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Price       *int64  `json:"price"`
		Currency    *string `json:"currency"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := valueobject.ParseUserID(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	productID, err := valueobject.ParseProductID(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id error"})
		return
	}

	err = h.updateProductUC.Execute(c.Request.Context(), usecase.UpdateProductCommand{
		UserID:      userID,
		ProductID:   productID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Currency:    req.Currency,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
