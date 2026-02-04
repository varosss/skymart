package event

import (
	"clirzy/product/internal/domain/valueobject"
	"time"
)

type ProductArchived struct {
	eventID    valueobject.EventID
	productID  valueobject.ProductID
	sellerID   valueobject.SellerID
	occurredAt time.Time
}

func NewProductArchived(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductArchived {
	return ProductArchived{
		eventID:    valueobject.NewEventID(),
		productID:  productID,
		sellerID:   sellerID,
		occurredAt: time.Now(),
	}
}

func (e ProductArchived) ID() string {
	return e.eventID.String()
}

func (ProductArchived) Type() string {
	return "product.archived"
}

func (e ProductArchived) AggregateID() string {
	return e.productID.String()
}

func (ProductArchived) AggregateType() string {
	return "product"
}

func (e ProductArchived) OccurredAt() time.Time {
	return e.occurredAt
}

func (e ProductArchived) ProductID() valueobject.ProductID {
	return e.productID
}

func (e ProductArchived) SellerID() valueobject.SellerID {
	return e.sellerID
}
