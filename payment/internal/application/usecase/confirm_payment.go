package usecase

import (
	aport "clirzy/payment/internal/application/port"
	"clirzy/payment/internal/domain"
	"clirzy/payment/internal/domain/port"
	"clirzy/payment/internal/domain/valueobject"
	"context"
)

type ConfirmPaymentCommand struct {
	PaymentID valueobject.PaymentID
}

type ConfirmPaymentUseCase struct {
	payments port.PaymentRepo
	bus      aport.EventBus
}

func NewConfirmPaymentUseCase(
	payments port.PaymentRepo,
	bus aport.EventBus,
) *ConfirmPaymentUseCase {
	return &ConfirmPaymentUseCase{
		payments: payments,
		bus:      bus,
	}
}

func (uc *ConfirmPaymentUseCase) Execute(
	ctx context.Context,
	cmd ConfirmPaymentCommand,
) error {

	payment, err := uc.payments.FindByID(ctx, cmd.PaymentID)
	if err != nil {
		return err
	}
	if payment == nil {
		return domain.ErrPaymentNotFound
	}

	if err := payment.MarkSucceeded(); err != nil {
		return err
	}

	if err := uc.payments.Save(ctx, payment); err != nil {
		return err
	}

	if err := uc.bus.Publish(ctx, payment.PullEvents()); err != nil {
		return err
	}

	return nil
}
