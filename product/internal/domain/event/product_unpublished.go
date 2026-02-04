package event

import (
	"clirzy/product/internal/domain/valueobject"
	"time"
)

type ProductUnpublished struct {
	eventID    valueobject.EventID
	productID  valueobject.ProductID
	sellerID   valueobject.SellerID
	occurredAt time.Time
}

func NewProductUnpublished(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductUnpublished {
	return ProductUnpublished{
		eventID:    valueobject.NewEventID(),
		productID:  productID,
		sellerID:   sellerID,
		occurredAt: time.Now(),
	}
}

func (e ProductUnpublished) ID() string {
	return e.eventID.String()
}

func (ProductUnpublished) Type() string {
	return "product.published"
}

func (e ProductUnpublished) AggregateID() string {
	return e.productID.String()
}

func (ProductUnpublished) AggregateType() string {
	return "product"
}

func (e ProductUnpublished) OccurredAt() time.Time {
	return e.occurredAt
}

func (e ProductUnpublished) ProductID() valueobject.ProductID {
	return e.productID
}

func (e ProductUnpublished) SellerID() valueobject.SellerID {
	return e.sellerID
}
