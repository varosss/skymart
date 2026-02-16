package yookassa

import (
	"clirzy/payment/internal/application/port"
	aport "clirzy/payment/internal/application/port"
	"context"
	"errors"
	"strconv"

	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

type YooKassaOptions struct {
	PaymentMethod string
	ReturnURL     string
	Description   string
	VatCode       int
}

type YooKassaProvider struct {
	client *yookassa.Client
}

func NewYooKassaProvider(shopID string, secretKey string) *YooKassaProvider {
	return &YooKassaProvider{
		client: yookassa.NewClient(shopID, secretKey),
	}
}

func (p *YooKassaProvider) CreatePayment(
	ctx context.Context,
	req port.CreatePaymentRequest,
) (*aport.CreatePaymentResponse, error) {
	paymentHandler := yookassa.NewPaymentHandler(p.client)

	money := &yoocommon.Amount{
		Value:    strconv.Itoa(int(req.Amount)),
		Currency: req.Currency,
	}

	receiptItem := yoocommon.Item{
		Description:    req.Description,
		Quantity:       "1.0",
		Amount:         money,
		VatCode:        1,
		PaymentSubject: "service",
		PaymentMode:    "full_payment",
	}

	yookassaPayment, err := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount:        money,
		PaymentMethod: yoopayment.PaymentMethodType(req.PaymentMethod),
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: req.ReturnURL,
		},
		Description: req.Description,
		Receipt: &yoopayment.Receipt{
			Customer: &yoocommon.Customer{Email: req.CustomerEmail},
			Items:    []*yoocommon.Item{&receiptItem},
		},
		Capture: true,
		Metadata: map[string]string{
			"payment_id": req.PaymentID,
		},
	})

	if err != nil {
		return nil, err
	}

	confirmation, ok := yookassaPayment.Confirmation.(map[string]any)
	if !ok {
		return nil, errors.New("unexpected confirmation type")
	}

	confirmationUrl, ok := confirmation["confirmation_url"].(string)
	if !ok {
		return nil, errors.New("invalid confirmation url")
	}

	return &aport.CreatePaymentResponse{
		Action: "redirect",
		ProviderData: map[string]string{
			"redirect_url": confirmationUrl,
		},
	}, nil
}

func (p *YooKassaProvider) CancelPayment(
	ctx context.Context,
	paymentID string,
) error {
	paymentHandler := yookassa.NewPaymentHandler(p.client)
	_, err := paymentHandler.CancelPayment(paymentID)

	return err
}
