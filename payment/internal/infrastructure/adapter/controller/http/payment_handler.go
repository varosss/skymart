package http

import (
	"clirzy/payment/internal/application/usecase"
	"clirzy/payment/internal/domain/valueobject"
	"net/http"

	"github.com/gin-gonic/gin"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	yoowebhook "github.com/rvinnie/yookassa-sdk-go/yookassa/webhook"
)

type PaymentHandler struct {
	confirmPaymentUC *usecase.ConfirmPaymentUseCase
	cancelPaymentUC  *usecase.CancelPaymentUseCase
}

func NewPaymentHandler(
	confirmPaymentUC *usecase.ConfirmPaymentUseCase,
	cancelPaymentUC *usecase.CancelPaymentUseCase,
) *PaymentHandler {
	return &PaymentHandler{
		confirmPaymentUC: confirmPaymentUC,
		cancelPaymentUC:  cancelPaymentUC,
	}
}

func (h *PaymentHandler) HandleYookassaWebhook(c *gin.Context) {
	var req yoowebhook.WebhookEvent[yoopayment.Payment]

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	metaRaw := req.Object.Metadata
	if metaRaw == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "metadata is empty"})
		return
	}

	meta, ok := metaRaw.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid metadata type"})
		return
	}

	rawPaymentID, ok := meta["payment_id"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_id not found"})
		return
	}

	paymentID, ok := rawPaymentID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_id is not string"})
		return
	}

	parsedPaymentID, err := valueobject.ParsePaymentID(paymentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch req.Object.Status {
	case yoopayment.Succeeded:
		err := h.confirmPaymentUC.Execute(c.Request.Context(), usecase.ConfirmPaymentCommand{
			PaymentID: parsedPaymentID,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	case yoopayment.Canceled:
		err := h.cancelPaymentUC.Execute(c.Request.Context(), usecase.CancelPaymentCommand{
			PaymentID: parsedPaymentID,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusOK)
}
