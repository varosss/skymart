package http

import (
	"clirzy/order/internal/application/usecase"
	"clirzy/order/internal/domain/valueobject"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	createOrderUC *usecase.CreateOrderUseCase
}

func NewOrderHandler(
	createOrderUC *usecase.CreateOrderUseCase,
) *OrderHandler {
	return &OrderHandler{
		createOrderUC: createOrderUC,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		BuyerID string `json:"buyer_id"`
		Items   []struct {
			SellerID  string `json:"seller_id"`
			ProductID string `json:"product_id"`
			Quantity  int64  `json:"quantity"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems := make([]usecase.CreateOrderItem, len(req.Items))
	for i, item := range req.Items {
		sellerID, err := valueobject.ParseSellerID(item.SellerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		productID, err := valueobject.ParseProductID(item.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderItems[i] = usecase.CreateOrderItem{
			SellerID:  sellerID,
			ProductID: productID,
			Quantity:  item.Quantity,
		}
	}

	buyerID, err := valueobject.ParseBuyerID(req.BuyerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := h.createOrderUC.Execute(c.Request.Context(), usecase.CreateOrderCommand{
		BuyerID: buyerID,
		Items:   orderItems,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order_id": orderID.String()})
}
