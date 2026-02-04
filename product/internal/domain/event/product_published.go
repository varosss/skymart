package event

import (
	"clirzy/product/internal/domain/valueobject"
	"time"
)

type ProductPublished struct {
	eventID    valueobject.EventID
	productID  valueobject.ProductID
	sellerID   valueobject.SellerID
	occurredAt time.Time
}

func NewProductPublished(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductPublished {
	return ProductPublished{
		eventID:    valueobject.NewEventID(),
		productID:  productID,
		sellerID:   sellerID,
		occurredAt: time.Now(),
	}
}

func (e ProductPublished) ID() string {
	return e.eventID.String()
}

func (ProductPublished) Type() string {
	return "product.published"
}

func (e ProductPublished) AggregateID() string {
	return e.productID.String()
}

func (ProductPublished) AggregateType() string {
	return "product"
}

func (e ProductPublished) OccurredAt() time.Time {
	return e.occurredAt
}

func (e ProductPublished) ProductID() valueobject.ProductID {
	return e.productID
}

func (e ProductPublished) SellerID() valueobject.SellerID {
	return e.sellerID
}
