package usecase

import (
	"clirzy/payment/internal/domain/port"
	"context"
)

type CreatePaymentCommand struct {
}

type CreatePaymentUseCase struct {
	paymentsRepo port.PaymentsRepo
}

func (uc *CreatePaymentUseCase) Execute(ctx context.Context, cmd CreatePaymentCommand) {
}
