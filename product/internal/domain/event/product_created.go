package event

import (
	"clirzy/product/internal/domain/valueobject"
	"time"
)

type ProductCreated struct {
	eventID    valueobject.EventID
	productID  valueobject.ProductID
	sellerID   valueobject.SellerID
	occurredAt time.Time
}

func NewProductCreated(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) *ProductCreated {
	return &ProductCreated{
		eventID:    valueobject.NewEventID(),
		productID:  productID,
		sellerID:   sellerID,
		occurredAt: time.Now(),
	}
}

func ProductCreatedFromPrimitives(
	eventID valueobject.EventID,
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	occurredAt time.Time,
) *ProductCreated {
	return &ProductCreated{
		eventID:    eventID,
		productID:  productID,
		sellerID:   sellerID,
		occurredAt: occurredAt,
	}
}

func (e *ProductCreated) ID() string {
	return e.eventID.String()
}

func (*ProductCreated) Type() string {
	return "product.created"
}

func (e *ProductCreated) AggregateID() string {
	return e.productID.String()
}

func (*ProductCreated) AggregateType() string {
	return "product"
}

func (e *ProductCreated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *ProductCreated) ProductID() valueobject.ProductID {
	return e.productID
}

func (e *ProductCreated) SellerID() valueobject.SellerID {
	return e.sellerID
}
