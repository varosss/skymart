package usecase

import (
	aport "clirzy/payment/internal/application/port"
	"clirzy/payment/internal/domain"
	"clirzy/payment/internal/domain/port"
	"clirzy/payment/internal/domain/valueobject"
	"context"
)

type CancelPaymentCommand struct {
	PaymentID string
}

type CancelPaymentUseCase struct {
	payments port.PaymentRepo
	provider aport.PaymentProvider
	bus      aport.EventBus
}

func NewCancelPaymentUseCase(
	payments port.PaymentRepo,
	provider aport.PaymentProvider,
	bus aport.EventBus,
) *CancelPaymentUseCase {
	return &CancelPaymentUseCase{
		payments: payments,
		provider: provider,
		bus:      bus,
	}
}

func (uc *CancelPaymentUseCase) Execute(
	ctx context.Context,
	cmd CancelPaymentCommand,
) error {

	paymentID, err := valueobject.ParsePaymentID(cmd.PaymentID)
	if err != nil {
		return domain.ErrInvalidPaymentID
	}

	payment, err := uc.payments.FindByID(ctx, paymentID)
	if err != nil {
		return err
	}
	if payment == nil {
		return domain.ErrPaymentNotFound
	}

	if err := payment.Cancel(); err != nil {
		return err
	}

	if err := uc.payments.Save(ctx, payment); err != nil {
		return err
	}

	if err := uc.provider.CancelPayment(ctx, paymentID.String()); err != nil {
		return err
	}

	if err := uc.bus.Publish(ctx, payment.PullEvents()); err != nil {
		return err
	}

	return nil
}
