package port

import (
	"clirzy/payment/internal/domain/entity"
	"context"
)

type PaymentsRepo interface {
	Save(ctx context.Context, payment *entity.Payment) error
}
