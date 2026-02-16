package event

import (
	"clirzy/product/internal/domain/valueobject"
	"time"
)

type ProductPriceUpdated struct {
	eventID    valueobject.EventID
	productID  valueobject.ProductID
	sellerID   valueobject.SellerID
	price      valueobject.Money
	occurredAt time.Time
}

func NewProductPriceUpdated(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	price valueobject.Money,
) *ProductPriceUpdated {
	return &ProductPriceUpdated{
		eventID:    valueobject.NewEventID(),
		productID:  productID,
		sellerID:   sellerID,
		price:      price,
		occurredAt: time.Now(),
	}
}

func ProductPriceUpdatedFromPrimitives(
	eventID valueobject.EventID,
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	price valueobject.Money,
	occurredAt time.Time,
) *ProductPriceUpdated {
	return &ProductPriceUpdated{
		eventID:    eventID,
		productID:  productID,
		sellerID:   sellerID,
		price:      price,
		occurredAt: occurredAt,
	}
}

func (e *ProductPriceUpdated) ID() string {
	return e.eventID.String()
}

func (*ProductPriceUpdated) Type() string {
	return "product.price_updated"
}

func (e *ProductPriceUpdated) AggregateID() string {
	return e.productID.String()
}

func (*ProductPriceUpdated) AggregateType() string {
	return "product"
}

func (e *ProductPriceUpdated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *ProductPriceUpdated) ProductID() valueobject.ProductID {
	return e.productID
}

func (e *ProductPriceUpdated) SellerID() valueobject.SellerID {
	return e.sellerID
}

func (e *ProductPriceUpdated) Price() valueobject.Money {
	return e.Price()
}
