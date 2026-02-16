package usecase

import (
	aport "clirzy/payment/internal/application/port"
	"clirzy/payment/internal/domain"
	"clirzy/payment/internal/domain/entity"
	"clirzy/payment/internal/domain/port"
	"clirzy/payment/internal/domain/valueobject"
	"context"
	"fmt"
)

type CreatePaymentCommand struct {
	InvoiceID     valueobject.InvoiceID
	CustomerEmail string
	PaymentMethod string
	ReturnURL     string
}

type PaymentActionType string

const (
	PaymentActionRedirect     PaymentActionType = "redirect"
	PaymentActionClientSecret PaymentActionType = "client_secret"
	PaymentActionNone         PaymentActionType = "none"
)

type PaymentAction struct {
	Type PaymentActionType
	Data map[string]string
}

type CreatePaymentResult struct {
	PaymentID string
	Action    PaymentAction
}

type CreatePaymentUseCase struct {
	payments port.PaymentRepo
	provider aport.PaymentProvider
	bus      aport.EventBus
	billing  aport.BillingGateway
}

func NewCreatePaymentUseCase(
	payments port.PaymentRepo,
	provider aport.PaymentProvider,
	bus aport.EventBus,
	billing aport.BillingGateway,
) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{
		payments: payments,
		provider: provider,
		bus:      bus,
		billing:  billing,
	}
}

func (uc *CreatePaymentUseCase) Execute(
	ctx context.Context,
	cmd CreatePaymentCommand,
) (*CreatePaymentResult, error) {
	payment, err := uc.payments.FindByInvoiceID(ctx, cmd.InvoiceID)
	if err != nil {
		return nil, err
	}
	if payment != nil {
		return nil, domain.ErrPaymentAlreadyExists
	}

	invoice, err := uc.billing.GetInvoiceByID(ctx, cmd.InvoiceID)
	if err != nil {
		return nil, fmt.Errorf("get invoice %s: %w", cmd.InvoiceID.String(), err)
	}

	amount, err := valueobject.NewMoney(invoice.Amount, invoice.Currency)
	if err != nil {
		return nil, err
	}

	payment, err = entity.NewPayment(
		cmd.InvoiceID,
		amount,
	)
	if err != nil {
		return nil, err
	}

	resp, err := uc.provider.CreatePayment(
		ctx,
		aport.CreatePaymentRequest{
			PaymentID:     payment.ID().String(),
			Amount:        amount.Amount(),
			Currency:      amount.Currency(),
			Description:   "Invoice #" + cmd.InvoiceID.String(),
			CustomerEmail: cmd.CustomerEmail,
			PaymentMethod: cmd.PaymentMethod,
			ReturnURL:     cmd.ReturnURL,
		},
	)

	if err != nil {
		return nil, err
	}

	if err := uc.payments.Save(ctx, payment); err != nil {
		return nil, err
	}

	if err := uc.bus.Publish(ctx, payment.PullEvents()); err != nil {
		return nil, err
	}

	return &CreatePaymentResult{
		PaymentID: payment.ID().String(),
		Action: PaymentAction{
			Type: PaymentActionType(resp.Action),
			Data: resp.ProviderData,
		},
	}, nil
}
